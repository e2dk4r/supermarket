package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/e2dk4r/supermarket/argon2id"
	"github.com/e2dk4r/supermarket/cockroachdb"
	"github.com/e2dk4r/supermarket/crypto"
	myHttp "github.com/e2dk4r/supermarket/http"
	myJwt "github.com/e2dk4r/supermarket/jwt"

	"github.com/jackc/pgx/v4"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)

	/* setup service dependencies */
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	/* setup services */
	productService := &cockroachdb.ProductService{
		Conn: conn,
	}

	orderService := &cockroachdb.OrderService{
		Conn: conn,
	}

	passwordService := &argon2id.PasswordService{
		Memory:      128 << 10, // 64mb
		Iterations:  4,
		Parallelism: 32,
		SaltLength:  32,
		KeyLength:   64,
	}

	userService := &cockroachdb.UserService{
		PasswordService: passwordService,
		Conn:            conn,
	}

	randomService := &crypto.RandomService{}

	jwtKey := os.Getenv("JWT_KEY")
	jwtSigner, err := jwt.NewSignerHS(jwt.HS256, []byte(jwtKey))
	if err != nil {
		log.Fatal(err)
	}
	jwtVerifier, err := jwt.NewVerifierHS(jwt.HS256, []byte(jwtKey))
	if err != nil {
		log.Fatal(err)
	}

	authService := &myJwt.AuthService{
		UserService:        userService,
		PasswordService:    passwordService,
		RandomService:      randomService,
		TokenValidDuration: 30 * time.Second,

		Signer:   jwtSigner,
		Verifier: jwtVerifier,
	}

	/* setup http */
	handler := myHttp.Handler{
		ProductService: productService,
		OrderService:   orderService,
		UserService:    userService,
		AuthService:    authService,
	}

	err = http.ListenAndServe(":8080", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
