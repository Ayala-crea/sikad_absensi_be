package handler

import (
	"fmt"
	"net/http"

	vercel "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := vercel.New()

	server.GET("/", func(context *vercel.Context) {
		context.JSON(200, vercel.H{
			"message": "hello go from vercel !!!!",
		})
	})
	server.GET("/hello", func(context *vercel.Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, vercel.H{
				"message": "name not found",
			})
		} else {
			context.JSON(200, vercel.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(context *vercel.Context) {
		context.JSON(400, vercel.H{
			"data": vercel.H{
				"id": context.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(context *vercel.Context) {
		context.JSON(200, vercel.H{
			"data": vercel.H{
				"url": context.Path,
			},
		})
	})
	server.Handle(w, r)
}
