package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main(){

	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://george:LAGos123@@localhost:5432/george"))
	if err != nil {
		fmt.Println("There is an error")
	}

	defer conn.Close(context.Background())

	var greeting string 
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Println("error is zero....")
	}
	fmt.Println(greeting)
}
