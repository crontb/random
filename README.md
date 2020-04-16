# Random

This is a basic REST API web service.

## Run app
With go runtime:
`go run cmd/main.go`

With Make (see Makefile for details):
`make all run`

With Docker:
`make docker`

The service will be running on port 8000.

## API endpoints
### /card
Generates 20 random integer numbers in 1-99 range with goroutines.

Methods: GET, OPTIONS

## Example
### Request
```
GET https://localhost:8000/card
```

### Response
```
Access-Control-Allow-Methods: GET,OPTIONS
Access-Control-Allow-Origin: *
Content-Length: 72
Content-Type: application/json
Date: Wed, 25 Mar 2020 22:10:26 GMT

{"result":[86,68,5,15,19,78,92,70,58,36,20,48,63,67,38,30,19,81,51,58]}
```

### CLI request
```
$ curl -k https://localhost:8000/card
```

### CLI response
```
{"result":[30,80,82,53,1,23,94,51,39,23,18,2,96,71,74,31,66,59,69,36]}
```

