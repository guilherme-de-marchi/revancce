package repository

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/jackc/pgx/v5"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/skip2/go-qrcode"
)

func PaymentWebhookOpenpixChargeCompleted(ctx context.Context, in model.PaymentWebhookOpenpixPostIn) error {
	var client model.ClientPostIn
	var clientID string
	clients, err := ClientGet(ctx, model.ClientGetIn{CPF: pkg.Varchar{Value: &in.CPF}})
	if err != nil && err != pgx.ErrNoRows {
		return pkg.Error(err)
	}

	if err == pgx.ErrNoRows || len(clients) == 0 {
		client = model.ClientPostIn{
			Name:  in.Name,
			Email: in.Email,
			CPF:   in.CPF,
			Phone: in.Phone,
		}

		v, err := ClientPost(ctx, client)
		if err != nil {
			return pkg.Error(err)
		}

		clientID = v.ID

	} else {
		clientID = clients[0].ID
	}

	outClientTicket, err := ClientTicketPost(ctx, model.ClientTicketPostIn{
		Client:      clientID,
		Batch:       in.Batch,
		Transaction: in.Transaction,
	})
	if err != nil {
		return pkg.Error(err)
	}

	// outEventBatch, err := EventBatchGet(
	// 	ctx,
	// 	model.EventBatchGetIn{ID: pkg.Varchar{Value: &in.Batch}},
	// )
	// if err != nil {
	// 	return pkg.Error(err)
	// }

	// if len(outEventBatch) == 0 {
	// 	return pkg.Error(ErrEventBatchNotFound)
	// }

	var eventName, ticketName string
	var batchPrice int
	row := pkg.Database.QueryRow(
		ctx,
		`
			select e.name, t.name, b.price
			from events_batches b
			join events_tickets t on t.id = b.ticket
			join events e on e.id = t.event
			where b.id = $1
		`,
		in.Batch,
	)
	if err := row.Scan(&eventName, &ticketName, &batchPrice); err != nil {
		return pkg.Error(err)
	}

	m := mail.NewSingleEmail(
		mail.NewEmail(pkg.SendGridPurchasesName, pkg.SendGridPurchasesEmail),
		"Ingresso registrado",
		// mail.NewEmail(in.Name, in.Email),
		mail.NewEmail(in.Name, pkg.SendGridPurchasesEmail),
		"",
		generatePaymentEmailTemplate(client.Name, eventName, ticketName, batchPrice),
	)

	qrcodePNG, err := qrcode.Encode(outClientTicket.ID, qrcode.Medium, 256)

	attachment := mail.NewAttachment()
	attachment.SetContent(base64.StdEncoding.EncodeToString(qrcodePNG))
	attachment.SetType("image/png")
	attachment.SetFilename("ticket.png")

	m.AddAttachment(attachment)
	resp, err := pkg.Mail.Send(m)
	if err != nil {
		return pkg.Error(err)
	}
	if resp.Body != "" {
		err = errors.New(resp.Body)
	}

	return pkg.Error(err)
}
