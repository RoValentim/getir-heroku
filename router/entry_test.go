package router

import (
        "testing"
        "net/http"
        "github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
        r := &Router{}

	r.Route(http.MethodPost, `/v1/database`,           nil)
	r.Route(http.MethodGet,  `/v1/memory`,             nil)
	r.Route(http.MethodPost, `/v1/memory`,             nil)

        notMatch := map[string]string(map[string]string(nil))
        match    := map[string]string(map[string]string{})

        msgNotMatch := "The route path should not match with wrong method"
        msgMatch    := "The route path should match"

        noErrorRequesting := "Can not have an error creating HTTP Request"

        request, err := http.NewRequest(http.MethodGet, "/v1/database", nil)
        assert.Equal(t, err, nil, noErrorRequesting)
        assert.Equal(t, notMatch, r.routes[0].Match(request), msgNotMatch)

        request, err = http.NewRequest(http.MethodPost, "/v1/database", nil)
        assert.Equal(t, err, nil, noErrorRequesting)
        assert.Equal(t, match, r.routes[0].Match(request), msgMatch)

        request, err = http.NewRequest(http.MethodGet, "/v1/memory?key=test", nil)
        assert.Equal(t, err, nil, noErrorRequesting)
        assert.Equal(t, map[string]string{"key":"test"}, r.routes[1].Match(request), "The route path should match with filter")

        request, err = http.NewRequest(http.MethodPost, "/v1/memory?key=test", nil)
        assert.Equal(t, err, nil, noErrorRequesting)
        assert.Equal(t, notMatch, r.routes[1].Match(request), msgNotMatch)

        request, err = http.NewRequest(http.MethodGet, "/v1/memory", nil)
        assert.Equal(t, err, nil, noErrorRequesting)
        assert.Equal(t, notMatch, r.routes[2].Match(request), msgNotMatch)

        request, err = http.NewRequest(http.MethodPost, "/v1/memory", nil)
        assert.Equal(t, err, nil, noErrorRequesting)
        assert.Equal(t, match, r.routes[2].Match(request), msgMatch)
}
