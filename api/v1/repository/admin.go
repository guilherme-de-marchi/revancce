package repository

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(ctx context.Context, req model.AdminLoginReq) (*model.AdminLoginResp, error) {
	row := pkg.Database.QueryRow(ctx, `
		select id, password_hash
		from admins
		where name=$1
	`, *req.Name)

	var id, passwordHash string
	if err := row.Scan(&id, &passwordHash); err != nil {
		return nil, pkg.Error(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(*req.Password)); err != nil {
		return nil, pkg.Error(err)
	}

	rawToken, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &pkg.RSA.PublicKey, []byte(id), nil)
	if err != nil {
		return nil, pkg.Error(err)
	}

	token := hex.EncodeToString(rawToken)

	exp := time.Hour
	if err := pkg.Memory.Set(ctx, "admin_session:"+id, token, exp).Err(); err != nil {
		return nil, pkg.Error(err)
	}

	return &model.AdminLoginResp{Token: token, ExpiresAt: time.Now().Add(exp)}, nil
}

func AdminRegister(ctx context.Context, req model.AdminRegisterReq) error {
	token, err := hex.DecodeString(req.HeaderAuthorization)
	if err != nil {
		return pkg.Error(err)
	}

	id, err := pkg.RSA.Decrypt(nil, token, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return pkg.Error(err)
	}

	originalToken, err := pkg.Memory.Get(ctx, "admin_session:"+string(id)).Result()
	if err != nil {
		return pkg.Error(err)
	}

	if req.HeaderAuthorization != originalToken {
		return pkg.Error(errors.New("session token mismatch"))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		return pkg.Error(err)
	}

	_, err = pkg.Database.Exec(ctx, `
		insert into admins (name, email, password_hash)
		values ($1, $2, $3)
	`, *req.Name, *req.Email, passwordHash)

	return pkg.Error(err)
}
