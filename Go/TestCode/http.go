// Monkey Patch Test
package TestCode

import (
	"bou.ke/monkey"
	e "errors"
	"net/http"
	"reflect"
)

var (
	ErrInvalidMethod = e.New("methods not allowed")
)

type Method string

const (
	GET   Method = "Get"
	PATCH Method = "Patch"
	POST  Method = "Post"
	PUT   Method = "Put"
)

func GetRequest(method string) error {
	var g *monkey.PatchGuard
	g = monkey.PatchInstanceMethod(
		reflect.TypeOf(http.DefaultClient),
		"Get",
		func(c *http.Client, url string) (*http.Response, error) {
			g.Unpatch()
			defer g.Restore()

			return c.Get(url)
		})

	if method == string(PUT) {
		return ErrInvalidMethod
	}

	return nil
}
