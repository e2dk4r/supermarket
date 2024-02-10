package http

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/e2dk4r/supermarket"
	"github.com/google/uuid"
)

func (h *Handler) OrderIndex(w http.ResponseWriter, r *http.Request) {
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

	orders, err := h.OrderService.Orders(page, perPage)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
		return
	}
	orderCount := len(orders)

	if orderCount == 0 {
		jsonFailResponse(w, http.StatusNotFound, fmt.Errorf("there is no product available"))
		return
	}

	type minifiedOrder struct {
		Id     string                  `json:"id,omitempty"`
		Status supermarket.OrderStatus `json:"status,omitempty"`
		Total  float32                 `json:"total,omitempty"`
	}

	minifiedOrders := make([]*minifiedOrder, 0, orderCount)

	for _, order := range orders {
		minifiedOrders = append(minifiedOrders, &minifiedOrder{
			Id:     order.Id,
			Status: order.Status,
			Total:  order.Total,
		})
	}

	jsonPaginatedResponse(w, page, perPage, 0, minifiedOrders)
}

func (h *Handler) OrderShow(w http.ResponseWriter, r *http.Request) {
	orderId := filepath.Base(r.URL.Path)

	if _, err := uuid.Parse(orderId); err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnprocessableEntity, fmt.Errorf("id must be in uuid format"))
		return
	}

	order, err := h.OrderService.Order(orderId)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
		return
	}

	jsonSuccessResponse(w, order)
}

func (h *Handler) OrderCreate(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("not implemented"))
}
