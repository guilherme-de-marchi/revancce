package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var (
	ErrWebhookFieldMalformatted = errors.New("webhook field malformatted")
)

type err struct {
	err   error
	Msg   string   `json:"msg"`
	Paths []string `json:"paths"`
}

func (e err) Error() string {
	d, _ := json.Marshal(e)
	return string(d)
}

func (e err) Is(target error) bool {
	return e.err == target
}

func Error(e error) error {
	if e == nil {
		return nil
	}

	_, filename, line, _ := runtime.Caller(1)
	here := fmt.Sprintf("%s:%v", filename, line)
	here = here[strings.Index(here, "/revancce/"):]

	newE, ok := e.(err)
	if !ok {
		return err{
			err:   e,
			Msg:   e.Error(),
			Paths: []string{here},
		}
	}

	newE.Paths = append(newE.Paths, here)
	return e
}
