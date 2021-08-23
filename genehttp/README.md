# Generic HTTP handler

With the generics coming to Go, it can make some generic handler that can't be handled by middleware way simpler and avoid lots of boilerplate.
The whole app can then focus on just the backend functions that are in the `func(Request) (Response, error)` form. 
And also, those piece of logic won't need generic code in them (hence the `api.go` is a normal Go1 file).

## Build and run Go (= 1.17)

```
go run -gcflags=-G=3 .
```

Then you can try those awesome features with curl:
```
curl localhost:8000/uppercase -d '{ "text": "this is lowercase" }'
{"Text":"THIS IS LOWERCASE"}
```

### Limitations

- you can't use generics on exported func
- you can't chain multiple generics function calls with generics parameters

## Build and run Go (< 1.17)

The simplest when you have docker is to use the [levonet/golang:go2go image](https://hub.docker.com/r/levonet/golang)

```
docker run --rm -v "$PWD:/go/src/app" -w /go/src/app levonet/golang:go2go go tool go2go build ; ./app
```

Then you can try those awesome features with curl:
```
curl localhost:8000/uppercase -d '{ "text": "this is lowercase" }'
{"Text":"THIS IS LOWERCASE"}
```
