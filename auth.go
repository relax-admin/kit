package kit

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSecret        = "JWT_SECRET"
)

const (
	JWT_KEY = "JWT_SECRET"
)

func init() {
	if s := os.Getenv(JWT_KEY); s != "" {
		jwtSecret = s
	}
}
func NewToken(m map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"iss": "account",
		"aud": "kitauth",
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(),
	}
	for k, v := range m {
		claims[k] = v
	}
	return jwt.NewWithClaims(jwtSigningMethod, claims).SignedString([]byte(jwtSecret))
}

func Renew(token string) (string, error) {
	claim, err := Extract(token)
	if err != nil {
		return "", err
	}
	claim["nbf"] = time.Now().Unix()
	claim["exp"] = time.Now().Add(time.Hour * 24 * 3).Unix()
	return jwt.NewWithClaims(jwtSigningMethod, claim).SignedString([]byte(jwtSecret))
}

func EditPayload(token string, m map[string]interface{}) (string, error) {
	claimInfo, err := Extract(token)
	if err != nil {
		return "", err
	}

	for k, v := range m {
		claimInfo[k] = v
	}

	return jwt.NewWithClaims(jwtSigningMethod, claimInfo).SignedString([]byte(jwtSecret))
}

func Extract(token string) (jwt.MapClaims, error) {
	if token == "" {
		return nil, fmt.Errorf("Required authorization token not found")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(jwtSecret), nil })
	if err != nil {
		return nil, fmt.Errorf("Error parsing token: %v", err)
	}

	if jwtSigningMethod != nil && jwtSigningMethod.Alg() != parsedToken.Header["alg"] {
		return nil, fmt.Errorf("Expected %s signing method but token specified %s",
			jwtSigningMethod.Alg(),
			parsedToken.Header["alg"])
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	claimInfo, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return claimInfo, nil
}
