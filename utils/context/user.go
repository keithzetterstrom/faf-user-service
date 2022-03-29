package context

import "context"

const UserIDHeader = "X-User-ID"

type userIDKey struct {}

func SetUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func GetUserID(ctx context.Context) *int64 {
	userID := ctx.Value(userIDKey{})
	if userID == nil {
		return nil
	}

	id := userID.(int64)
	return &id
}
