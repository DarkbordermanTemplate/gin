package tests

import (
	"practice/endpoints"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func Test_fruit_by_path_success(t *testing.T) {
	Setup()
	r := gofight.New()
	r.GET("/fruit/apple").
		SetDebug(true).
		Run(endpoints.Create_router(), func(rp gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, 200, rp.Code)
			assert.Equal(t, `{"count":1,"name":"apple"}`, rp.Body.String())
		})
}
func Test_fruit_by_path_fail(t *testing.T) {
	Setup()
	r := gofight.New()
	r.GET("/fruit/nonexist").
		SetDebug(true).
		Run(endpoints.Create_router(), func(rp gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, 400, rp.Code)
			assert.Equal(t, `Bad Request`, rp.Body.String())
		})
}

func Test_fruit_by_query_success(t *testing.T) {
	Setup()
	r := gofight.New()
	r.GET("/fruit").
		SetDebug(true).
		SetQuery(gofight.H{
			"name": "apple",
		}).
		Run(endpoints.Create_router(), func(rp gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, 200, rp.Code)
			assert.Equal(t, `{"count":1,"name":"apple"}`, rp.Body.String())
		})
}

func Test_fruit_by_query_fail(t *testing.T) {
	Setup()
	r := gofight.New()
	r.GET("/fruit/nonexist").
		SetDebug(true).
		Run(endpoints.Create_router(), func(rp gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, 400, rp.Code)
			assert.Equal(t, `Bad Request`, rp.Body.String())
		})
}
