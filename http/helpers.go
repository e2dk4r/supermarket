package http

import (
	"errors"
	"net/http"
	"strconv"
)

func getQueryPageParam(r *http.Request) (int, error) {
	page, err := getQueryIntParam(r, "page", 1)
	if err != nil {
		return 0, err
	}
	if page <= 0 {
		return 0, errors.New("page must be bigger than 0")
	}

	return page, nil
}

func getQueryPerPageParam(r *http.Request) (int, error) {
	page, err := getQueryIntParam(r, "perPage", 1)
	if err != nil {
		return 0, err
	}
	if page <= 0 {
		return 0, errors.New("per page must be bigger than 0")
	}

	return page, nil
}

func getQueryIntParam(r *http.Request, param string, def int) (int, error) {
	input := r.URL.Query().Get(param)
	if input == "" {
		return def, nil
	}

	parsed, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return parsed, nil
}
