package jwt

import (
	"errors"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/e2dk4r/supermarket"
)

type AuthService struct {
	UserService     supermarket.UserService
	PasswordService supermarket.PasswordService
	RandomService   supermarket.RandomService

	Signer             jwt.Signer
	Verifier           jwt.Verifier
	TokenValidDuration time.Duration
}

func (as *AuthService) CreateToken(u *supermarket.User) (string, error) {
	// fetch user from storage
	user, err := as.UserService.User(u.Username)
	if err != nil {
		return "", err
	}

	// verify hashed password matches with user
	match, err := as.PasswordService.Verify(u.Password, user.Password)
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.New("passwords does not match")
	}

	// creating jw token
	builder := jwt.NewBuilder(as.Signer)

	id, err := as.RandomService.GenerateString(32)
	if err != nil {
		return "", err
	}

	token, err := builder.Build(&jwt.RegisteredClaims{
		ID:       id,
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Issuer:   user.Id,
	})

	if err != nil {
		return "", err
	}

	return token.String(), nil
}

func (as *AuthService) VerifyToken(key string) error {
	token, err := jwt.Parse([]byte(key), as.Verifier)

	if err != nil {
		return err
	}

	var claims jwt.RegisteredClaims
	err = token.DecodeClaims(&claims)
	if err != nil {
		return err
	}

	if claims.IssuedAt == nil {
		return errors.New("token must have iat claim")
	}
	if claims.IssuedAt.Add(as.TokenValidDuration).Before(time.Now()) {
		return errors.New("expired token")
	}

	return nil
}
