package application

import (
	"context"
	"errors"
	"net/http"
)

func GetAppFromContext(ctx context.Context) (Application, error) {
	if app, ok := ctx.Value(ContextApp).(Application); ok {
		return app, nil
	}

	return nil, errors.New("cannot get app container from request")
}

func GetAppFromRequest(r *http.Request) (Application, error) {
	return GetAppFromContext(r.Context())
}
