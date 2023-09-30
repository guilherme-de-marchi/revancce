package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrWebhookFieldMalformatted = errors.New("webhook field malformatted")
)

type Err struct {
	Err   error
	Msg   string   `json:"msg"`
	Paths []string `json:"paths"`
}

func (e Err) Error() string {
	d, _ := json.Marshal(e)
	return string(d)
}

func (e Err) Is(target error) bool {
	return e.Err == target
}

func Error(e error) error {
	if e == nil {
		return nil
	}

	_, filename, line, _ := runtime.Caller(1)
	here := fmt.Sprintf("%s:%v", filename, line)
	here = here[strings.Index(here, "/revancce/"):]

	newE, ok := e.(Err)
	if !ok {
		return Err{
			Err:   e,
			Msg:   e.Error(),
			Paths: []string{here},
		}
	}

	newE.Paths = append(newE.Paths, here)
	return e
}

func ErrorMsg(msg string) any {
	return map[string]string{"error": msg}
}

func ErrorToPgError(err error) *pgconn.PgError {
	pkgErr, ok := err.(Err)
	if !ok {
		return nil
	}

	pgErr, ok := pkgErr.Err.(*pgconn.PgError)
	if !ok {
		return nil
	}

	return pgErr
}
