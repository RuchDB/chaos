package main

import (
	"github.com/savsgio/atreugo/v11"
)

var (
	corsAllowHeaders     = "*"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:5000",
	}
	server := atreugo.New(config)

	server.UseBefore(func(ctx *atreugo.RequestCtx) error {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		return ctx.Next()
	})

	server.GET("/stats", countStatsHandler)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

func countStatsHandler(ctx *atreugo.RequestCtx) error {
	stats := countStats()

	total := 0
	for _, value := range stats {
		total += value
	}

	return ctx.JSONResponse(atreugo.JSON{
		"stats": stats,
		"total": total,
	}, 200)
}
