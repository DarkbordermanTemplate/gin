package tests

import (
	"practice/endpoints"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func Test_health(t *testing.T) {
	Setup()
	r := gofight.New()
	r.GET("/health").
		SetDebug(true).
		Run(endpoints.Create_router(), func(rp gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, 200, rp.Code)
			assert.Equal(t, `OK`, rp.Body.String())
		})
}
