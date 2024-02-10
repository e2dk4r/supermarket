package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/e2dk4r/supermarket/cockroachdb"
	myHttp "github.com/e2dk4r/supermarket/http"

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

	/* setup http */
	handler := myHttp.Handler{
		ProductService: productService,
		OrderService:   orderService,
	}

	err = http.ListenAndServe(":8080", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
