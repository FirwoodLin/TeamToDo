package response

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Success bool        `json:"success"`
	Hint    string      `json:"hint"`
	Data    interface{} `json:"data"`
}

func MakeResponse(success bool, hint string, data interface{}) response {
	return response{
		Success: success,
		Hint:    hint,
		Data:    data,
	}
}

func MakeSucceedResponse(data interface{}) response {
	return MakeResponse(true, "", data)
}

func MakeFailedResponse(hint string) response {
	return MakeResponse(false, hint, nil)
}

func MakeFailedResponse2(ctx *gin.Context, status int, hint string) {
	ctx.JSON(status, MakeResponse(false, hint, nil))
	// return MakeResponse(false, hint, nil)
}
