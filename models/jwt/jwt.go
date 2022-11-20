package jwt

// import (
// 	"be-golang-project/consts"
// 	"be-golang-project/models/interface_"
// 	"be-golang-project/models/payload"
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// type JWTMaker struct {
// 	secretKey string
// }

// func NewJWTMaker(secretKey string) (interface_.Maker, error) {
// 	if len(secretKey) < consts.MinSecretKeySize {
// 		return nil, fmt.Errorf("Invalid key size: must be at least %d characters", int(consts.MinSecretKeySize))
// 	}

// 	return &JWTMaker{secretKey}, nil
// }

// func (maker *JWTMaker) CreateToken(userID int, duration time.Duration) (string, error) {
// 	payload, err := payload.NewPayload(userID, duration)
// 	if err != nil {
// 		return "", err
// 	}

// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

// 	return jwtToken.SignedString([]byte(maker.secretKey))
// }

// func (maker *JWTMaker) VerifyToken(token string) (*payload.Payload, error) {
// 	var (
// 		errExpiredToken = errors.New(consts.ErrExpiredToken)
// 		errInvalidToken = errors.New(consts.ErrInvalidToken)
// 	)

// 	keyFunc := func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)
// 		if !ok {
// 			return nil, errInvalidToken
// 		}

// 		return []byte(maker.secretKey), nil
// 	}

// 	jwtToken, err := jwt.ParseWithClaims(token, &payload.Payload{}, keyFunc)
// 	if err != nil {
// 		verr, _ := err.(*jwt.ValidationError)
// 		if strings.Contains(verr.Inner.Error(), consts.ErrExpiredToken) {
// 			return nil, errExpiredToken
// 		}

// 		return nil, errInvalidToken
// 	}

// 	payload, ok := jwtToken.Claims.(*payload.Payload)
// 	if !ok {
// 		return nil, errInvalidToken
// 	}

// 	return payload, nil
// }
