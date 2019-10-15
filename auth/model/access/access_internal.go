package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"micro-service/basic/config"
	"time"
)

func (s *service) createTokenClaims(subject *Subject) (m *jwt.StandardClaims, err error) {
	now := time.Now()
	m = &jwt.StandardClaims{
		ExpiresAt: now.Add(tokenExpiredDate).Unix(),
		NotBefore: now.Unix(),
		Id:        subject.ID,
		IssuedAt:  now.Unix(),
		Issuer:    "book.micro.mu",
		Subject:   subject.ID,
	}

	return
}

func (s *service) SaveTokenToCache(subject *Subject, val string) (err error) {
	if err = ca.Set(tokenIDKeyPrefix+subject.ID, val, tokenExpiredDate).Err(); err != nil {
		return fmt.Errorf("[saveTokenToCache] Save Token Fail :" + err.Error())
	}
	return
}

func (s *service) delTokenFromCache(subject *Subject) (err error) {
	if err = ca.Del(tokenIDKeyPrefix + subject.ID).Err(); err != nil {
		return fmt.Errorf("[delTokenFromCache] Del Token Fail: " + err.Error())
	}
	return
}

func (s *service) getTokenFromCache(subject *Subject) (token string, err error) {
	tokenCached, err := ca.Get(tokenIDKeyPrefix + subject.ID).Result()
	if err != nil {
		return token, fmt.Errorf("[getTokenfromCache] token not exists %s", err)
	}
	return string(tokenCached), nil
}

func (s *service) parseToken(tk string) (c *jwt.StandardClaims, err error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (i interface{}, e error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token format: %v", token.Header["alg"])
		}
		return []byte(config.GetJwtConfig().GetSecretKey()), nil
	})

	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("[parseToken] Expired Token,err: %s", err)
			default:
				break
			}
			break
		default:
			break
		}

		return nil, fmt.Errorf("[parseToken] invalid Token,err: %s", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("[parseToken] invalid Token")
	}

	return mapClaimToJwClaim(claims), nil
}

func mapClaimToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {
	jC := &jwt.StandardClaims{Subject: claims["sub"].(string)}
	return jC
}
