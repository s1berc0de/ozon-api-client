package report

import (
	v1 "github.com/s1berc0de/ozon-api-client/ozon/report/v1"
	"net/http"
)

func New(
	h *http.Client,
	uri string,
) *Report {
	return &Report{
		v1: v1.New(h, uri+"/v1/report"),
	}
}

type Report struct {
	v1 *v1.Report
}

func (c *Report) V1() *v1.Report {
	return c.v1
}
