package main

import (
	"contex/ctxx"
	"context"
	"fmt"
)

type contextKey string
type UserID int

const (
	userIDKey contextKey = "userID"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, UserID(123))
	fmt.Println(ctx.Value(userIDKey).(UserID))

	ctx2 := context.Background()
	ctx2 = ctxx.WithSingleton(ctx2, UserID(123))
	fmt.Println(ctxx.Singleton[UserID](ctx2))

}
