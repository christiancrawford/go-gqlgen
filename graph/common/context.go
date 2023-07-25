package common

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

// This code is creating a custom context to pass the database connection through the GraphQL request context:

// CustomContext struct holds the database connection
type CustomContext struct {
	Database *gorm.DB
}

// customContextKey is used to store the context in the request context
var customContextKey string = "CUSTOM_CONTEXT"

// CreateContext middleware attaches the CustomContext to each request
func CreateContext(args *CustomContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customContext := &CustomContext{
			Database: args.Database,
		}
		requestWithCtx := r.WithContext(context.WithValue(r.Context(), customContextKey, customContext))
		next.ServeHTTP(w, requestWithCtx)
	})
}

// GetContext retrieves the CustomContext from the request context
func GetContext(ctx context.Context) *CustomContext {
	customContext, ok := ctx.Value(customContextKey).(*CustomContext)
	if !ok {
		return nil
	}
	return customContext
}
