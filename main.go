package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"log"
	"microTest/proto"
)

type Say struct {
	Client hello.HelloService
}

func (s *Say) Hello(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	log.Print("Received Say.Say API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.hello", "Name cannot be blank")
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": "hello, " + name.Values[0],
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.hello"),
	)
	service.Init()
	hello.RegisterHelloHandler(service.Server(), &Say{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
