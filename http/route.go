package http

import (
	"strings"

	"github.com/google/uuid"
)

// RouteMatch is matching to see if actual requested url path matches
// route url. Dynamic routes are matched with "." character
// Returns true if it matches, false otherwise
func RouteMatch(requestedUrl, routeUrl string) bool {
	// is it dynamic route
	index := strings.Index(routeUrl, ".")

	// if it is not dynamic route, url must match exactly
	if index < 0 {
		return requestedUrl == routeUrl
	}

	lSplit := strings.Split(requestedUrl, "/")
	rSplit := strings.Split(routeUrl, "/")

	if len(lSplit) != len(rSplit) {
		return false
	}

	for i := 0; i < len(lSplit); i++ {
		dynamicRoute := strings.HasPrefix(rSplit[i], ".")

		matchAny := dynamicRoute && rSplit[i] == "."
		if matchAny {
			continue
		}

		matchUuid := dynamicRoute && rSplit[i] == ".uuid"
		if matchUuid {
			_, err := uuid.Parse(lSplit[i])
			if err != nil {
				return false
			}
			continue
		}

		if lSplit[i] != rSplit[i] {
			return false
		}
	}

	return true
}
