package model

import "time"

type PaymentWebhookOpenpixPostReq struct {
	Event  string `json:"event"`
	Charge struct {
		Customer          interface{}         `json:"customer"`
		Value             int                 `json:"value"`
		Comment           string              `json:"comment"`
		Identifier        string              `json:"identifier"`
		PaymentLinkID     string              `json:"paymentLinkID"`
		TransactionID     string              `json:"transactionID"`
		Status            string              `json:"status"`
		AdditionalInfo    []map[string]string `json:"additionalInfo"`
		Discount          int                 `json:"discount"`
		ValueWithDiscount int                 `json:"valueWithDiscount"`
		ExpiresDate       time.Time           `json:"expiresDate"`
		Type              string              `json:"type"`
		CorrelationID     string              `json:"correlationID"`
		CreatedAt         time.Time           `json:"createdAt"`
		UpdatedAt         time.Time           `json:"updatedAt"`
		PaidAt            time.Time           `json:"paidAt"`
		Payer             struct {
			Name  string `json:"name"`
			TaxID struct {
				TaxID string `json:"taxID"`
				Type  string `json:"type"`
			} `json:"taxID"`
			Email         string `json:"email"`
			Phone         string `json:"phone"`
			CorrelationID string `json:"correlationID"`
		} `json:"payer"`
		BrCode         string `json:"brCode"`
		ExpiresIn      int    `json:"expiresIn"`
		PixKey         string `json:"pixKey"`
		PaymentLinkURL string `json:"paymentLinkUrl"`
		QrCodeImage    string `json:"qrCodeImage"`
		GlobalID       string `json:"globalID"`
	} `json:"charge"`
	Pix struct {
		Customer struct {
			Name  string `json:"name"`
			TaxID struct {
				TaxID string `json:"taxID"`
				Type  string `json:"type"`
			} `json:"taxID"`
			Email         string `json:"email"`
			Phone         string `json:"phone"`
			CorrelationID string `json:"correlationID"`
		} `json:"customer"`
		Payer struct {
			Name  string `json:"name"`
			TaxID struct {
				TaxID string `json:"taxID"`
				Type  string `json:"type"`
			} `json:"taxID"`
			Email         string `json:"email"`
			Phone         string `json:"phone"`
			CorrelationID string `json:"correlationID"`
		} `json:"payer"`
		Charge struct {
			Customer          interface{}   `json:"customer"`
			Value             int           `json:"value"`
			Comment           string        `json:"comment"`
			Identifier        string        `json:"identifier"`
			PaymentLinkID     string        `json:"paymentLinkID"`
			TransactionID     string        `json:"transactionID"`
			Status            string        `json:"status"`
			AdditionalInfo    []interface{} `json:"additionalInfo"`
			Discount          int           `json:"discount"`
			ValueWithDiscount int           `json:"valueWithDiscount"`
			ExpiresDate       time.Time     `json:"expiresDate"`
			Type              string        `json:"type"`
			CorrelationID     string        `json:"correlationID"`
			CreatedAt         time.Time     `json:"createdAt"`
			UpdatedAt         time.Time     `json:"updatedAt"`
			PaidAt            time.Time     `json:"paidAt"`
			Payer             struct {
				Name  string `json:"name"`
				TaxID struct {
					TaxID string `json:"taxID"`
					Type  string `json:"type"`
				} `json:"taxID"`
				Email         string `json:"email"`
				Phone         string `json:"phone"`
				CorrelationID string `json:"correlationID"`
			} `json:"payer"`
			BrCode         string `json:"brCode"`
			ExpiresIn      int    `json:"expiresIn"`
			PixKey         string `json:"pixKey"`
			PaymentLinkURL string `json:"paymentLinkUrl"`
			QrCodeImage    string `json:"qrCodeImage"`
			GlobalID       string `json:"globalID"`
		} `json:"charge"`
		Value         int       `json:"value"`
		Time          time.Time `json:"time"`
		EndToEndID    string    `json:"endToEndId"`
		TransactionID string    `json:"transactionID"`
		InfoPagador   string    `json:"infoPagador"`
		Type          string    `json:"type"`
		CreatedAt     time.Time `json:"createdAt"`
		GlobalID      string    `json:"globalID"`
	} `json:"pix"`
	Company struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		TaxID string `json:"taxID"`
	} `json:"company"`
	Account struct {
	} `json:"account"`
}

type PaymentWebhookOpenpixPostIn struct {
	Type        string
	Name        string
	CPF         string
	Email       string
	Phone       string
	Transaction string
	Batch       string
}
