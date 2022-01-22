package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func recovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			p := recover()
			if p != nil {
				log.Printf("panic in web server %s", p)
				fmt.Println(string(debug.Stack()))
				context.String(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		context.Next()
	}
}
