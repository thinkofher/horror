package horror

import (
	"context"
	"net/http"
)

func defaultInternalServerErrorHandler(e error, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(e.Error()))
	return
}

type horrorCtxKey int

const internalServerErrorHandlerKey = horrorCtxKey(0)

type errorHandlerFunc func(error, http.ResponseWriter, *http.Request)

func getInternalServerErrorHandler(ctx context.Context) errorHandlerFunc {
	f, ok := ctx.Value(internalServerErrorHandlerKey).(errorHandlerFunc)
	if !ok {
		return defaultInternalServerErrorHandler
	}
	return f
}

func setInternalServerErrorHandler(ctx context.Context,
	f errorHandlerFunc) context.Context {
	return context.WithValue(
		ctx, internalServerErrorHandlerKey, f,
	)
}

func RegisterInternalServerErrorHandler(
	f func(error, http.ResponseWriter, *http.Request),
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newCtx := setInternalServerErrorHandler(
				r.Context(),
				errorHandlerFunc(f),
			)
			next.ServeHTTP(w, r.WithContext(newCtx))
			return
		})
	}
}
