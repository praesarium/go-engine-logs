package logs

import (
	"github.com/praesarium/go-engine/engine"
	"github.com/praesarium/go-colors/colors"
	"log"
	"os"
	"time"
)

func MiddlewareLogs() engine.Middleware {

	format := time.Now().Format("2006/01/02 - 15:04:05")
	logger := log.New(os.Stdout, "", 0)

	logger.Printf("%s[SERVER]%s %v - server application started",
		colors.BG_MAGENTA,
		colors.EFT_RESET,
		format,
	)

	return func(ctx *engine.Context) {

		// determine total time of request
		timeStart := time.Now()
		ctx.NextMiddleware()
		timeAfter := time.Now()

		method := ctx.Request.Method
		status := ctx.Writer.Code()

		logger.Printf("%s[REQUEST]%s %v - %s| %3d |%s %s\t%s %s\t%010v - %12v - %s",
			colors.BG_GREEN,
			colors.EFT_RESET,
			timeAfter.Format("2006/01/02 - 15:04:05"),
			httpStatusColor(status),
			status,
			colors.EFT_RESET,
			httpMethodColor(ctx.Request.Method),
			colors.EFT_RESET,
			method,
			timeAfter.Sub(timeStart),
			ctx.Request.RemoteAddr,
			ctx.Request.RequestURI,
		)
	}
}

func httpMethodColor(method string) string {
	switch {
	case method == "OPTIONS":
		{
			return colors.BG_MAGENTA
		}
	case method == "HEAD":
		{
			return colors.BG_MAGENTA
		}
	case method == "GET":
		{
			return colors.BG_BLUE
		}
	case method == "PUT":
		{
			return colors.BG_YELLOW
		}
	case method == "POST":
		{
			return colors.BG_CYAN
		}
	case method == "DELETE":
		{
			return colors.BG_RED
		}
	case method == "PATCH":
		{
			return colors.BG_GREEN
		}
	default:
		{
			return colors.EFT_RESET
		}
	}
}

func httpStatusColor(code int) string {
	switch {
	case code >= 200 && code <= 299:
		{
			return colors.BG_GREEN
		}
	case code >= 300 && code <= 399:
		{
			return colors.BG_YELLOW
		}
	case code >= 400 && code <= 499:
		{
			return colors.BG_MAGENTA
		}
	default:
		{
			return colors.BG_RED
		}
	}
}
