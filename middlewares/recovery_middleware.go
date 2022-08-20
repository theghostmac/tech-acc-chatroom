package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.Logger
}

func NewMiddleware(logger *zap.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}

func (rm *Middleware) RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {

				// if you're implementing any logging lib (e.g zap, logrus)
				// then use as required, and return the error stack trace, example below for zap
				// handler.logger.Debug("error string", zap.Any("stack", string(debug.stack())))

				log.Print(err)
				rm.logger.Debug("error trace", zap.Any("stack", string(debug.Stack())))
				w.WriteHeader(http.StatusOK)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
