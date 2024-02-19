package middleware

import (
	"errors"

	tele "go.mamad.dev/telebot"
)

//
//// AutoRespond returns a middleware that automatically responds
//// to every callback.
//func AutoRespond() tele.MiddlewareFunc {
//	return func(next tele.HandlerFunc) tele.HandlerFunc {
//		return func(c tele.Context) error {
//			if c.Callback() != nil {
//				defer func(c tele.Context, resp ...*tele.CallbackResponse) {
//					err := c.Respond(resp...)
//					if err != nil {
//						fmt.Printf("[telebot] Error, failed to AutoRespond: %s", err.Error())
//					}
//				}(c)
//			}
//			return next(c)
//		}
//	}
//}
//
//// IgnoreVia returns a middleware that ignores all the
//// "sent via" messages.
//func IgnoreVia() tele.MiddlewareFunc {
//	return func(next tele.HandlerFunc) tele.HandlerFunc {
//		return func(c tele.Context) error {
//			if msg := c.AccessibleMessage(); msg != nil && msg.Via != nil {
//				return nil
//			}
//			return next(c)
//		}
//	}
//}

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
