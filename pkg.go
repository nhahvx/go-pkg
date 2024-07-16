package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/nhahvx/go-pkg/src/http/middleware"
)

type IPkg interface {
	AuthMiddleware(jwtKey string) gin.HandlerFunc
}

type Pkg struct {
}

func New() IPkg {
	return Pkg{}
}

func (p Pkg) AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return middleware.AuthMiddleware(jwtKey)
}
