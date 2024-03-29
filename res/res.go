package res

import (
	"log"
	"net/http"
	error "websiteGin/error"

	"github.com/gin-gonic/gin"
)

type response struct {
	Rtn    int         `json:"rtn"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

/*
res.SendOK(c, data)
res.SendParamErr(c, error.ErrUserLogin, "user id: xxxx")
res.SendParamError(c, err)
*/

func ok(data interface{}) *response {
	return &response{
		Rtn:    0,
		ErrMsg: "",
		Data:   data,
	}
}

func fail(code int, msg string) *response {
	return &response{
		Rtn:    code,
		ErrMsg: msg,
		Data:   nil,
	}
}

func paramErr(code int, msg string) *response {
	if code == 0 {
		return fail(code, msg)
	} else {
		return fail(error.CodeParam, msg)
	}
}

func serverErr(code int, msg string) *response {
	if code == 0 {
		return fail(code, msg)
	} else {
		return fail(error.CodeServer, msg)
	}
}

func SendOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ok(data))
}

func SendParamError(ctx *gin.Context, errCode int, msg string) {
	if errCode == 0 {
		errCode = error.CodeParam
	}
	log.Println(errCode, msg)
	// todo: must not return msg to front end user.
	ctx.JSON(http.StatusBadRequest, paramErr(errCode, msg))
}

func SendServerError(ctx *gin.Context, errCode int, msg string) {
	if errCode == 0 {
		errCode = error.CodeServer
	}
	log.Println(errCode, msg)
	ctx.JSON(http.StatusInternalServerError, serverErr(errCode, msg))
}
