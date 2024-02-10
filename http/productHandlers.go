package http

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
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

func (h *Handler) ProductCreate(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("not implemented"))
}
