
package router

import (
        "testing"
        "net/http"
        "github.com/stretchr/testify/assert"
 )

func TestRoute(t *testing.T) {
        msg := "The value shuold exist and the method should be the same"

        r := &Router{}

        r.Route(http.MethodPost, `/v1/database`, nil)
	r.Route(http.MethodGet,  `/v1/memory`,   nil)
	r.Route(http.MethodPost, `/v1/memory`,   nil)

        assert.Equal(t, r.routes[0].Method, http.MethodPost, msg)
        assert.Equal(t, r.routes[1].Method, http.MethodGet,  msg)
        assert.Equal(t, r.routes[2].Method, http.MethodPost, msg)
}
