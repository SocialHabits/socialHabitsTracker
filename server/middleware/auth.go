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
	RoleName   string
	IsLoggedIn bool
}

// SetToken method to write cookie
func (access *CookieAccess) SetToken(token string) {
	http.SetCookie(access.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
	})
}

type CookieContent struct {
	UserId   int
	RoleName string
}

func extractUserIdAndRoleName(ctx *gin.Context) (*CookieContent, error) {
	c, err := ctx.Request.Cookie("jwt")
	if err != nil {
		return nil, errors.New("There is no token in cookies")
	}

	claims, err := util.ValidateIdToken(c.Value)
	if err != nil {
		return nil, err
	}

	return &CookieContent{UserId: claims.UserID, RoleName: claims.RoleName}, nil
}

func setValInCtx(ctx *gin.Context, val interface{}) {
	newCtx := context.WithValue(ctx.Request.Context(), cookieAccessKeyCtx, val)
	ctx.Request = ctx.Request.WithContext(newCtx)
}

func GetValFromCtx(ctx context.Context) *CookieAccess {
	raw := ctx.Value(cookieAccessKeyCtx).(*CookieAccess)

	return raw
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookieA := CookieAccess{
			Writer: ctx.Writer,
		}

		// &cookieA is a pointer so any changes in future is changing cookieA is context
		setValInCtx(ctx, &cookieA)

		user, err := extractUserIdAndRoleName(ctx)

		if err != nil {
			cookieA.IsLoggedIn = false
			cookieA.RoleName = ""
			ctx.Next()
			return
		}

		cookieA.UserId = uint64(user.UserId)
		cookieA.IsLoggedIn = true
		cookieA.RoleName = user.RoleName

		// calling the actual resolver
		ctx.Next()
		// here will execute after resolver and all other middlewares was called
		// so &cookieA is safe from garbage collector
	}
}

func GetCookieAccess(ctx context.Context) *CookieAccess {
	return ctx.Value(cookieAccessKeyCtx).(*CookieAccess)
}
