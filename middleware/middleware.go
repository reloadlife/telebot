package middleware

import (
	"errors"
	tele "go.mamad.dev/telebot"
)

// Recover returns a middleware that recovers a panic happened in
// the handler.
func Recover(onError ...func(error)) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			var f func(error)
			if len(onError) > 0 {
				f = onError[0]
			} else {
				f = func(err error) {
					c.Bot().OnError(err, nil)
				}
			}

			defer func() {
				if r := recover(); r != nil {
					switch er := r.(type) {
					case error:
						f(er)
					case string:
						f(errors.New(er))
					}
				}
			}()

			return next(c)
		}
	}
}
