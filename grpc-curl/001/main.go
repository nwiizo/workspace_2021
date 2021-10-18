package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	pb _
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

const (
	address = "xx.xx.xx.xx:443"
)

func get_jwt() string {
	// Set up a connection to the server.
	var tlsConf tls.Config
	creds := credentials.NewTLS(&tlsConf)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewXXXXXXXXXxClient(conn)

	// Contact the server and print out its response.
	//XXX := "12345"
	//YYY := "13"

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	r, err := c.SignIn(ctx, &pb.Xxxxxx{XxxxxXx: XXX, Yyyyyyy: YYY})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	return r.GetJwtToken()
}

func main() {
	fmt.Println("SignIn Request:", get_jwt())
}
