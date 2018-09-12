# Default API

This example makes use of the default micro API request handler.

The api expects you use the [api.Request/Response](https://github.com/micro/go-api/blob/master/proto/api.proto) protos.

The default micro api request handler gives you full control over the http request and response while still leveraging RPC and 
any transport plugins that use other protocols beyond http in your stack such as grpc, nats, kafka.

## Usage

## Ps :run local add prefix :   MICRO_REGISTRY=mdns

Run the micro API

```
micro api
```

## local 
```
MICRO_REGISTRY=mdns micro api
```

## service : example
Run this example

```
go run default.go
```

## local 
```
MICRO_REGISTRY=mdns go run default.go
```

Make a GET request to /example/call which will call go.micro.api.example Example.Call

```
curl "http://localhost:8080/example/call?name=john"
```

Make a POST request to /example/foo/bar which will call go.micro.api.example Foo.Bar

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```



## service : newapi
Run this newapi

```
go run newapi.go
```

## local 
```
MICRO_REGISTRY=mdns go run newapi.go
```

Make a GET request to /newapi/call which will call go.micro.api.newapi Example.Call

```
curl "http://localhost:8080/newapi/call?name=john"
```

Make a POST request to /newapi/foo/bar which will call go.micro.api.newapi newapi2.Bar

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/newapi/newapi2/bar
```
#
