package jwt

import (
	"fmt"
	"food-delivery/component/tokenprovider"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"tokenPayload"`
	jwt.StandardClaims
}

// Generate ...
func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &tokenprovider.Token{
		Token:   myToken,
		Created: time.Now(),
		Expiry:  expiry,
	}, nil
}

// Validate ...
func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {
	fmt.Println("myToken: ", myToken)
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		fmt.Println("check 1")
		return nil, tokenprovider.ErrInvalidToken
	}

	if !res.Valid {
		fmt.Println("check 2")

		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		fmt.Println("check 3")

		return nil, tokenprovider.ErrInvalidToken
	}

	// Return the token
	return &claims.Payload, nil
}
