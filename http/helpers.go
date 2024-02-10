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
	perPage, err := getQueryIntParam(r, "perPage", 25)
	if err != nil {
		return 0, err
	}
	if perPage <= 0 {
		return 0, errors.New("per page must be bigger than 0")
	}
	if perPage > 500_000 {
		return 0, errors.New("per page must be smaller than 500_000")
	}

	var valid = []int{5, 25, 50, 250, 500}
	ok := false
	for _, i := range valid {

		if i == perPage {
			ok = true
			break
		}
	}
	if !ok {
		return 0, errors.New("per page is invalid")
	}

	return perPage, nil
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
