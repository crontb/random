# Random

This is a basic REST API web service.

## Run app
`go run cmd/main.go`

## API endpoints
### /card
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


