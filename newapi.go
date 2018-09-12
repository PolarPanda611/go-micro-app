package main

import (
	"encoding/json"
	"log"
	"strings"

	proto2 "examples/api/default/proto2"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"

	"context"
)

type Newapi struct{}

type Newapi2 struct{}

// Example.Call is a method which will be served by http request /example/call
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /example/call goes to go.micro.api.example Example.Call
func (e *Newapi) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Example.Call request")

	// parse values from the get request
	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.newapi", "no content")
	}

	// set response status
	rsp.StatusCode = 200

	// respond with some json
	b, _ := json.Marshal(map[string]string{
		"message": "got your request " + strings.Join(name.Values, " "),
	})

	// set json body
	rsp.Body = string(b)

	return nil
}

// Foo.Bar is a method which will be served by http request /example/foo/bar
// Because Foo is not the same as the service name it is mapped beyond /example/
func (f *Newapi2) Bar(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Newapi2.Bar request")

	// check method
	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.newapi", "require post")
	}

	// let's make sure we get json
	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.newapi", "need content-type")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.newapi", "expect application/json")
	}

	// parse body
	var body map[string]interface{}
	err := json.Unmarshal([]byte(req.Body), &body)

	if err != nil {
		return errors.BadRequest("go.micro.api.newapi", "json unmarchal err"+string(err.Error()))
	}
	rsp.StatusCode = 201
	b, _ := json.Marshal(map[string]string{
		"message": "Post Successfully",
	})
	rsp.Body = string(b)
	// do something with parsed body

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.newapi"),
	)

	service.Init()

	// register example handler
	proto2.RegisterNewapiHandler(service.Server(), new(Newapi))

	// register foo handler
	proto2.RegisterNewapi2Handler(service.Server(), new(Newapi2))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
