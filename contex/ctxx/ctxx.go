package ctxx

import "context"

type ctxKey[T any] struct{}

// WithSingleton attaches val to the context as a singleton.
func WithSingleton[T any](ctx context.Context, val T) context.Context {
	return context.WithValue(ctx, ctxKey[T]{}, val)
}

// Singleton returns the single value T attached to the context.
// If there is no value attached, the zero value is returned.
func Singleton[T any](ctx context.Context) T {
	val, _ := ctx.Value(ctxKey[T]{}).(T)
	return val
}
