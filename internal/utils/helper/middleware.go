package helper

import (
	"context"

	"github.com/google/uuid"
	constanta "github.com/rchmachina/rach-fw/internal/const"
)

func GetRequestID(ctx context.Context) string {
	key := constanta.RequestIDKey

	if v := ctx.Value(key); v != nil {
		if reqID, ok := v.(string); ok {
			return reqID
		}
	}

	return ""
}

func SetRequestID(ctx context.Context, id string) context.Context {
	reqID := uuid.NewString()
	return context.WithValue(ctx, constanta.RequestIDKey, reqID)
}
