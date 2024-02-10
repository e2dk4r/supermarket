package http

import (
	"fmt"
	"net/http"

	"github.com/e2dk4r/supermarket"
)

type Handler struct {
	ProductService  supermarket.ProductService
	OrderService    supermarket.OrderService
	UserService     supermarket.UserService
	AuthService     supermarket.AuthService
	PasswordService supermarket.PasswordService
	RandomService   supermarket.RandomService
	DbErrorService  supermarket.DbErrorService
}

// middlewareFunc gets http writer and http request for inspecting
// and modifying. returns true if next middleware should be processed,
// false if it should be stopped
type middlewareFunc func(http.ResponseWriter, *http.Request) bool

type route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	Middlewares []middlewareFunc
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	globalMiddlewares := []middlewareFunc{
		h.MSecurityHeaders,
	}

	routes := []route{
		// users
		{Method: http.MethodPost, Path: "/signup", Handler: h.Signup},
		{Method: http.MethodPost, Path: "/login", Handler: h.Login},
		{Method: http.MethodGet, Path: "/token", Handler: h.Token},

		// products
		{Method: http.MethodGet, Path: "/product/.", Handler: h.ProductShow},
		{Method: http.MethodGet, Path: "/products", Handler: h.ProductIndex},
		{Method: http.MethodPost, Path: "/product/create", Handler: h.ProductCreate, Middlewares: []middlewareFunc{h.MRequireAuth}},
		{Method: http.MethodPost, Path: "/product/delete", Handler: h.ProductDelete},

		// orders
		{Method: http.MethodGet, Path: "/order/.", Handler: h.OrderShow},
		{Method: http.MethodGet, Path: "/orders", Handler: h.OrderIndex},
		{Method: http.MethodPost, Path: "/order/create", Handler: h.OrderCreate},
		{Method: http.MethodPost, Path: "/order/delete", Handler: h.OrderCreate},
	}

	for _, route := range routes {
		match := r.Method == route.Method && RouteMatch(r.URL.Path, route.Path)
		if !match {
			continue
		}

		for _, middleware := range globalMiddlewares {
			if !middleware(w, r) {
				return
			}
		}
		for _, middleware := range route.Middlewares {
			next := middleware(w, r)
			if !next {
				return
			}
		}

		route.Handler(w, r)
		return
	}

	h.NotFound(w, r)
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusNotFound, fmt.Errorf("not found"))
}
