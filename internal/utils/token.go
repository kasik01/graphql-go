package utils

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func mergeMap(map1, map2 map[string]interface{}) map[string]interface{} {
	for key, value := range map1 {
		map2[key] = value
	}
	return map2
}

func genClaims(rawClaims map[string]interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	for key, value := range rawClaims {
		claims[key] = value
	}
	return claims
}

func CreateToken(customClaims map[string]interface{}, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Second * time.Duration(expiry)).Unix()
	defaultClaims := map[string]interface{}{
		"exp": exp,
	}
	claims := genClaims(mergeMap(defaultClaims, customClaims))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func VerifyToken(tokenString string, secret string) (any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		hmacSampleSecret := []byte(secret)
		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["payload"], nil
	} else {
		return nil, err
	}
}
