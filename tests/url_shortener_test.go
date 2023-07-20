package tests

import (
	"URL_Shortener/internal/http-server/handlers/url/save"
	"URL_Shortener/internal/lib/random"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
	"net/url"
	"testing"
)

const (
	host = "localhost:1234"
)

func TestURLShortenerServe(t *testing.T) {

	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.Default(t, u.String())

	e.POST("/url").
		WithJSON(save.Request{
			URL:   gofakeit.URL(),
			Alias: random.NewRandomString(10),
		}).
		WithBasicAuth("myuser", "mypass").
		Expect().
		Status(200).
		JSON().Object().
		ContainsKey("alias")
}
