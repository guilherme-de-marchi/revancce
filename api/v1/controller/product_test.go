package controller

import (
	"net/http"
	"testing"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

func TestGetProduct_Success(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, pkg.ServerDomain+"/api/v1/product/1", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("wrong status code: %v", resp.StatusCode)
	}
}

func TestPurchaseProduct_Success(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, pkg.ServerDomain+"/api/v1/product/1/purchase", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("wrong status code: %v", resp.StatusCode)
	}
}
