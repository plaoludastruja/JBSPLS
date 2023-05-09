package http

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Context *gin.Context
}

func sendResponse(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, data)
}

func (gin *Gin) GenericResponse(code int, message string, data interface{}) {
	gin.Context.JSON(code, data)
}

func (gin *Gin) OK(data interface{}) {
	sendResponse(gin.Context, 200, "OK", data)
}

func (gin *Gin) Created(data interface{}) {
	sendResponse(gin.Context, 201, "Created", data)
}

func (gin *Gin) Accepted(data interface{}) {
	sendResponse(gin.Context, 202, "Accepted", data)
}

func (gin *Gin) NoContent(data interface{}) {
	sendResponse(gin.Context, 204, "No Content", data)
}

func (gin *Gin) BadRequest(data interface{}) {
	sendResponse(gin.Context, 400, "Bad Request", data)
}

func (gin *Gin) Unauthorized(data interface{}) {
	sendResponse(gin.Context, 401, "Unauthorized", data)
}

func (gin *Gin) NotFound(data interface{}) {
	sendResponse(gin.Context, 404, "Not Found", data)
}
