package repository

import (
	"context"
	"time"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(ctx context.Context, in model.AdminLoginIn) (*model.AdminLoginOut, error) {
	row := pkg.Database.QueryRow(
		ctx,
		`
			select id, password_hash
			from admins
			where name=$1
		`,
		in.Name,
	)

	var authorization, passwordHash string
	if err := row.Scan(&authorization, &passwordHash); err != nil {
		return nil, pkg.Error(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(in.Password)); err != nil {
		return nil, pkg.Error(err)
	}

	token, exp, err := pkg.NewAdminSession(ctx, authorization)
	if err != nil {
		return nil, pkg.Error(err)
	}

	return &model.AdminLoginOut{Token: token, ExpiresAt: time.Now().Add(exp)}, nil
}

func AdminRegister(ctx context.Context, in model.AdminRegisterIn) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return pkg.Error(err)
	}

	_, err = pkg.Database.Exec(
		ctx,
		`
		insert into admins (name, email, password_hash, created_by)
		values ($1, $2, $3, $4)
		`,
		in.Name, in.Email, passwordHash, in.AdminID,
	)

	return pkg.Error(err)
}
