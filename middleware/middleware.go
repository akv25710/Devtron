package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func RecoverMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					if r == http.ErrAbortHandler {
						panic(r)
					}
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					var stack []byte
					var length int

					stack = make([]byte, 1<<10)
					length = runtime.Stack(stack, true)
					stack = stack[:length]

					msg := fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:length])
					logrus.Error(msg)
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
