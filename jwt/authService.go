package jwt

import (
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/e2dk4r/supermarket"
)

type AuthService struct {
	UserService     supermarket.UserService
	PasswordService supermarket.PasswordService
	RandomService   supermarket.RandomService
	TimeService     supermarket.TimeService

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
		return "", supermarket.ErrPasswordNotMatch
	}

	// creating jw token
	builder := jwt.NewBuilder(as.Signer)

	id, err := as.RandomService.GenerateString(32)
	if err != nil {
		return "", err
	}

	token, err := builder.Build(&jwt.RegisteredClaims{
		ID:       id,
		IssuedAt: jwt.NewNumericDate(as.TimeService.Now()),
		Subject:  user.Id,
	})

	if err != nil {
		return "", err
	}

	return token.String(), nil
}

func (as *AuthService) CreateAnonToken() (string, error) {
	id, err := as.RandomService.GenerateString(32)
	if err != nil {
		return "", err
	}

	token, err := jwt.NewBuilder(as.Signer).Build(&jwt.RegisteredClaims{
		ID:       id,
		IssuedAt: jwt.NewNumericDate(as.TimeService.Now()),
	})

	if err != nil {
		return "", err
	}

	return token.String(), nil
}

func (as *AuthService) VerifyToken(key string) error {
	return as.verifyToken(key, false)
}

func (as *AuthService) VerifyAnonToken(key string) error {
	return as.verifyToken(key, true)
}

func (as *AuthService) verifyToken(key string, anon bool) error {
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
		return supermarket.ErrTokenIatClaim
	}
	if claims.IssuedAt.Add(as.TokenValidDuration).Before(as.TimeService.Now()) {
		return supermarket.ErrTokenExpired
	}
	if !anon && claims.Subject == "" {
		return supermarket.ErrTokenSubjectClaim
	}

	return nil
}
