package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/e2dk4r/supermarket"
)

func (h *Handler) ProductIndex(w http.ResponseWriter, r *http.Request) {
	page, err := getQueryPageParam(r)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("page not recognized"))
		return
	}
	perPage, err := getQueryPerPageParam(r)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("per page not recognized"))
		return
	}

	products, err := h.ProductService.Products(page, perPage)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	if len(products) == 0 {
		jsonFailResponse(w, http.StatusNotFound, fmt.Errorf("there is no product available"))
		return
	}

	jsonSuccessResponse(w, products)
}

func (h *Handler) ProductShow(w http.ResponseWriter, r *http.Request) {
	productId := filepath.Base(r.URL.Path)

	product, err := h.ProductService.Product(productId)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	jsonSuccessResponse(w, product)
}

const productRequestMaxBytes = 512 << 10 // 512kb

func (h *Handler) ProductCreate(w http.ResponseWriter, r *http.Request) {
	p := supermarket.Product{}
	defer r.Body.Close()

	// decode json
	err := json.NewDecoder(io.LimitReader(r.Body, productRequestMaxBytes)).Decode(&p)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("product not specified"))
		return
	}

	// validate product
	err = p.Validate()
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = h.ProductService.CreateProduct(&p)
	if err != nil {
		log.Print(err)

		if h.ProductService.IsDuplicateError(err) {
			jsonFailResponse(w, http.StatusBadRequest, fmt.Errorf("duplicate product"))
			return
		}

		jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("cannot create product"))
		return
	}

	jsonSuccessResponse(w, p)
}
