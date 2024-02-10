package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/e2dk4r/supermarket"
	"github.com/google/uuid"
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
		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
		return
	}

	if len(products) == 0 {
		jsonFailResponse(w, http.StatusNotFound, fmt.Errorf("there is no product available"))
		return
	}

	jsonPaginatedResponse(w, page, perPage, 0, products)
}

func (h *Handler) ProductShow(w http.ResponseWriter, r *http.Request) {
	productId := filepath.Base(r.URL.Path)

	if _, err := uuid.Parse(productId); err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("id must be in uuid format"))
		return
	}

	product, err := h.ProductService.Product(productId)
	if err != nil {
		log.Print(err)

		if h.ProductService.IsNotFoundError(err) {
			jsonFailResponse(w, http.StatusNotFound, supermarket.ErrNotFound)
			return
		}

		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
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

		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
		return
	}

	jsonSuccessResponse(w, p)
}

func (h *Handler) ProductDelete(w http.ResponseWriter, r *http.Request) {
	var product supermarket.Product
	defer r.Body.Close()

	limited := io.LimitReader(r.Body, productRequestMaxBytes)
	err := json.NewDecoder(limited).Decode(&product)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("cannot parse product id"))
		return
	}

	if product.Id == "" {
		jsonFailResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("id required"))
		return
	}

	_, err = uuid.Parse(product.Id)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("product id must be in uuid format"))
		return
	}

	deleted, err := h.ProductService.DeleteProduct(product.Id)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
		return
	}
	if !deleted {
		log.Print(err)
		jsonFailResponse(w, http.StatusNotFound, supermarket.ErrNotFound)
		return
	}

	jsonSuccessResponseMessage(w, "product deleted")
}
