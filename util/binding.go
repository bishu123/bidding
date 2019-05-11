package util

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func BindRequestFromJson(context *gin.Context, response interface{}) bool {
    context.Header("Content-Type", "application/json; charset=utf-8")
    if err := context.Bind(response); err!= nil{
        context.JSON(http.StatusBadRequest,nil)
        return false
    }
    return true
}
