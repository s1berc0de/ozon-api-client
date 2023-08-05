package logger

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

var _ http.RoundTripper = (*RoundTripper)(nil)

func NewRoundTripper(r http.RoundTripper) *RoundTripper {
	return &RoundTripper{r: r}
}

type RoundTripper struct {
	r http.RoundTripper
}

func (h RoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	d, _ := httputil.DumpRequestOut(r, true)
	fmt.Println(d) // log

	resp, err := h.r.RoundTrip(r)

	dump, e := httputil.DumpResponse(resp, true)
	if e == nil {
		fmt.Println(dump) // log
	}

	return resp, err
}
