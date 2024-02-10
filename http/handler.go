package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/e2dk4r/supermarket"
)

type Handler struct {
	ProductService  supermarket.ProductService
	OrderService    supermarket.OrderService
	UserService     supermarket.UserService
	AuthService     supermarket.AuthService
	PasswordService supermarket.PasswordService
	RandomService   supermarket.RandomService
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

		// products
		{Method: http.MethodGet, Path: "/product/{productId}", Handler: h.ProductShow},
		{Method: http.MethodGet, Path: "/products", Handler: h.ProductIndex},
		{Method: http.MethodPost, Path: "/product/create", Handler: h.ProductCreate, Middlewares: []middlewareFunc{h.MRequireAuth}},
		{Method: http.MethodPost, Path: "/product/delete", Handler: h.ProductDelete},

		// orders
		{Method: http.MethodGet, Path: "/order/{orderId}", Handler: h.OrderIndex},
		{Method: http.MethodPost, Path: "/order/create", Handler: h.OrderCreate},
		{Method: http.MethodGet, Path: "/orders", Handler: h.OrderIndex},
	}

	for _, route := range routes {
		match := r.Method == route.Method && routeMatch(r.URL.Path, route.Path)
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

func routeMatch(requestedUrl, routeUrl string) bool {
	if !strings.Contains(routeUrl, "{") && requestedUrl == routeUrl {
		return true
	}

	// is it dynamic route
	index := strings.Index(routeUrl, "{")
	if index < 0 {
		return false
	}

	// check if matches dynamic route
	if !strings.HasPrefix(requestedUrl, routeUrl[:index]) {
		return false
	}

	// only one dynamic variable supported
	if strings.Contains(requestedUrl[index:], "/") {
		return false
	}

	return true
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusNotFound, fmt.Errorf("not found"))
}
