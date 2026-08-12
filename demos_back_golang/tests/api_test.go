package tests

import (
	"net/url"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

const (
	host = "localhost:8081"
)

func TestAPI_HealthCheck(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host: host,
	}
	e := httpexpect.Default(t, u.String())

	e.GET("/v1/health").
	Expect().
	Status(200).
	JSON().
	Object().ContainsKey("Message")
}