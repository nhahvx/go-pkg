package security

import "github.com/golang-jwt/jwt/v5"

func GenerateJwtToken(claims jwt.Claims, jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
