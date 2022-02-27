package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
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
