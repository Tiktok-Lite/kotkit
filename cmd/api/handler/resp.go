package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func ResponseSuccess(c *app.RequestContext, response interface{}) {
	c.JSON(http.StatusOK, response)
}

func ResponseError(c *app.RequestContext, httpErrorCode int, response interface{}) {
	c.JSON(httpErrorCode, response)
}
