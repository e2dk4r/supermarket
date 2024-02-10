package http_test

import (
	"testing"

	"github.com/e2dk4r/supermarket/http"
)

type RouteMatchTest struct {
	RequestedUrl string
	RouteUrl     string
	IsMatch      bool
}

func TestRouteMatch(t *testing.T) {
	testCases := []RouteMatchTest{
		{RequestedUrl: "/", RouteUrl: "/", IsMatch: true},
		{RequestedUrl: "/foo", RouteUrl: "/", IsMatch: false},
		{RequestedUrl: "/foo", RouteUrl: "/foo", IsMatch: true},
		{RequestedUrl: "/foo.txt", RouteUrl: "/foo.txt", IsMatch: true},
		{RequestedUrl: "/foo/boo", RouteUrl: "/foo/.", IsMatch: true},
		{RequestedUrl: "/foo/123", RouteUrl: "/foo/.", IsMatch: true},
		{RequestedUrl: "/foo/boo", RouteUrl: "/foo/./.", IsMatch: false},
		{RequestedUrl: "/foo/foo", RouteUrl: "/./boo", IsMatch: false},
		{RequestedUrl: "/foo/boo", RouteUrl: "/./boo", IsMatch: true},
		{RequestedUrl: "/foo/boo/hoo", RouteUrl: "/foo/./hoo", IsMatch: true},
		{RequestedUrl: "/foo/boo/hoo", RouteUrl: "/foo/./foo", IsMatch: false},
		{RequestedUrl: "/foo/boo/hoo", RouteUrl: "/foo/./.", IsMatch: true},
	}

	for _, testCase := range testCases {
		actual := http.RouteMatch(testCase.RequestedUrl, testCase.RouteUrl)
		if actual != testCase.IsMatch {
			expected := "match"
			if !testCase.IsMatch {
				expected = "not match"
			}

			t.Errorf("expected route to %v. requestUrl: %v routeUrl: %v",
				expected,
				testCase.RequestedUrl,
				testCase.RouteUrl,
			)
		}

	}
}
