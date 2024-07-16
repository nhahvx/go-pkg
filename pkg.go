package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nhahvx/go-pkg/src/http/middleware"
	"github.com/nhahvx/go-pkg/src/security"
	"github.com/nhahvx/go-pkg/src/xml_helper"
)

type IPkg interface {

	// Http
	AuthMiddleware(jwtKey string) gin.HandlerFunc

	// XML
	ConvertStructToXML(v interface{}) (string, error)
	ConvertXMLToJSON(xmlString string) (string, error)

	// Sercurity
	GenerateJwtToken(claims jwt.Claims, jwtKey string) (string, error)
}

type Pkg struct {
}

func (p Pkg) GenerateJwtToken(claims jwt.Claims, jwtKey string) (string, error) {
	return security.GenerateJwtToken(claims, jwtKey)
}

func (p Pkg) ConvertXMLToJSON(xmlString string) (string, error) {
	return xml_helper.ConvertXMLToJSON(xmlString)
}

func (p Pkg) ConvertStructToXML(v interface{}) (string, error) {
	return xml_helper.ConvertStructToXML(v)
}

func (p Pkg) AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return middleware.AuthMiddleware(jwtKey)
}

func New() IPkg {
	return Pkg{}
}
