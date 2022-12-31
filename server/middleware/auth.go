package middleware

import (
	"context"
	"errors"
	"github.com/AntonioTrupac/socialHabitsTracker/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const cookieAccessKeyCtx = "cookieAccess"

type CookieAccess struct {
	Writer     http.ResponseWriter
	UserId     uint64
	IsLoggedIn bool
}

// method to write cookie
func (access *CookieAccess) SetToken(token string) {
	http.SetCookie(access.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
	})
}

func extractUserId(ctx *gin.Context) (int, error) {
	c, err := ctx.Request.Cookie("jwt")
	if err != nil {
		return 0, errors.New("There is no token in cookies")
	}

	claims, err := util.ValidateIdToken(c.Value)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func setValInCtx(ctx *gin.Context, val interface{}) {
	newCtx := context.WithValue(ctx.Request.Context(), cookieAccessKeyCtx, val)
	ctx.Request = ctx.Request.WithContext(newCtx)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookieA := CookieAccess{
			Writer: ctx.Writer,
		}

		// &cookieA is a pointer so any changes in future is changing cookieA is context
		setValInCtx(ctx, &cookieA)

		userID, err := extractUserId(ctx)

		if err != nil {
			cookieA.IsLoggedIn = false
			ctx.Next()
			return
		}

		cookieA.UserId = uint64(userID)
		cookieA.IsLoggedIn = true

		// calling the actual resolver
		ctx.Next()
		// here will execute after resolver and all other middlewares was called
		// so &cookieA is safe from garbage collector
	}
}

func GetCookieAccess(ctx context.Context) *CookieAccess {
	return ctx.Value(cookieAccessKeyCtx).(*CookieAccess)
}
