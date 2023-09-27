package pkg

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

func NewAdminSession(ctx context.Context, id string) (string, time.Duration, error) {
	rawToken, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &RSA.PublicKey, []byte(id), nil)
	if err != nil {
		return "", 0, Error(err)
	}

	token := hex.EncodeToString(rawToken)
	exp := time.Hour
	err = Memory.Set(ctx, "admin_session:"+id, token, exp).Err()

	return token, exp, Error(err)
}

func GetAdminSession(ctx context.Context, authorization string) (string, error) {
	token, err := hex.DecodeString(authorization)
	if err != nil {
		return "", Error(err)
	}

	id, err := RSA.Decrypt(nil, token, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return "", Error(err)
	}

	originalToken, err := Memory.Get(ctx, "admin_session:"+string(id)).Result()
	if err != nil {
		return "", Error(err)
	}

	if authorization != originalToken {
		return "", Error(errors.New("session token mismatch"))
	}

	return string(id), nil
}
