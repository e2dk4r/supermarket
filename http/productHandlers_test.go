package http_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/e2dk4r/supermarket"
	"github.com/e2dk4r/supermarket/http"
	"github.com/e2dk4r/supermarket/mock"
)

func TestProductShowWhenProductExists(t *testing.T) {
	// create mock ProductService
	ps := &mock.ProductService{
		ProductFn: func(id string) (*supermarket.Product, error) {
			return &supermarket.Product{
				Id:    "2f0495b9-099e-4c3f-9803-a4b8e32448a5",
				Name:  "Onion",
				Price: 3.50,
			}, nil
		},
	}

	h := http.Handler{
		ProductService: ps,
	}

	req := httptest.NewRequest("GET", "/product/2f0495b9-099e-4c3f-9803-a4b8e32448a5", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	resp := w.Result()

	// assert
	if !ps.ProductInvoked() {
		t.Errorf("did not call Product() from product service")
	}

	if resp.StatusCode != 200 {
		t.Errorf("must respond with http 200")
	}

	var model supermarket.Product
	_ = json.NewDecoder(resp.Body).Decode(&model)

	if model.Id != "2f0495b9-099e-4c3f-9803-a4b8e32448a5" {
		t.Errorf("id is not same")
	}

	if model.Name != "Onion" {
		t.Errorf("name is not same")
	}

	if model.Price != 3.5 {
		t.Errorf("price is not same")
	}
}

func TestProductShowWhenProductDoesntExist(t *testing.T) {
	// create mock ProductService
	ps := &mock.ProductService{
		ProductFn: func(id string) (*supermarket.Product, error) {
			return nil, fmt.Errorf("no row exists")
		},
		IsNotFoundErrorFn: func(err error) bool {
			return true
		},
	}

	h := http.Handler{
		ProductService: ps,
	}

	req := httptest.NewRequest("GET", "/product/2f0495b9-099e-4c3f-9803-a4b8e32448a5", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	resp := w.Result()

	// assert
	if !ps.ProductInvoked() {
		t.Errorf("did not call Product() from product service")
	}

	if !ps.IsNotFoundErrorInvoked() {
		t.Errorf("did not call IsNotFoundError() from product service")
	}

	if resp.StatusCode != 404 {
		t.Errorf("must respond with http 404")
	}

	var model struct {
		Error string `json:"error"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&model)

	if model.Error != "not found" {
		t.Errorf("responded with wrong error message")
	}
}

func TestProductShowWhenDatabaseError(t *testing.T) {
	// create mock ProductService
	ps := &mock.ProductService{
		ProductFn: func(id string) (*supermarket.Product, error) {
			return nil, fmt.Errorf("timeout")
		},
		IsNotFoundErrorFn: func(err error) bool {
			return false
		},
	}

	h := http.Handler{
		ProductService: ps,
	}

	req := httptest.NewRequest("GET", "/product/2f0495b9-099e-4c3f-9803-a4b8e32448a5", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	resp := w.Result()

	// assert
	if !ps.ProductInvoked() {
		t.Errorf("did not call Product() from product service")
	}

	if !ps.IsNotFoundErrorInvoked() {
		t.Errorf("did not call IsNotFoundError() from product service")
	}

	if resp.StatusCode != 500 {
		t.Errorf("must respond with http 500")
	}

	var model struct {
		Error string `json:"error"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&model)

	if model.Error != "internal server error" {
		t.Errorf("responded with wrong error message")
	}
}

func TestProductShowWhenProductIdIsNotUUID(t *testing.T) {
	// create mock ProductService
	ps := &mock.ProductService{
		ProductFn: func(id string) (*supermarket.Product, error) {
			return nil, fmt.Errorf("database error")
		},
	}

	h := http.Handler{
		ProductService: ps,
	}

	req := httptest.NewRequest("GET", "/product/xx", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	resp := w.Result()

	// assert
	if ps.ProductInvoked() {
		t.Errorf("must not call Product() from product service")
	}

	if ps.IsNotFoundErrorInvoked() {
		t.Errorf("must not call IsNotFoundError() from product service")
	}

	if resp.StatusCode != 422 {
		t.Errorf("must respond with http 422")
	}

	var model struct {
		Error string `json:"error"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&model)

	if model.Error != "id must be in uuid format" {
		t.Errorf("responded with wrong error message")
	}
}
