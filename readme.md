# Responder

## Index
- [Installation](#instalacion)
- [Usage](#usage)

# Installation

```sh
go get -u gitlab.com/digitalhouse-dev/libraries-go/util.git
```

# Usage

## Response Package

```go
// Import package
import (
    "gitlab.com/digitalhouse-dev/libraries-go/util.git/response"
)
// Example struct
type Car struct {
    Brand string `json:"brand"`
}
```
### Success response
```go
car := &Car{"Ferrari"}
res := response.Success("success", car, nil, headers.New())
```
Output
```json
{
    "message": "success",
    "code": 200,
    "data" {
        "brand": "Ferrari"
    }
    "meta": null
}
```
### Response with meta
```go
car := &Car{"Ferrari"}
meta := response.NewMeta(1, 10, 150)
res := response.Success("success", car, meta, nil)
```
Output
```json
{
    "message": "success",
    "code": 200,
    "data" {
        "brand": "Ferrari"
    }
    "meta": {
        "page": 1,
        "per_page": 10,
        "page_count": 15,
        "total_count": 100,
    }
}
```
### Bad Request response
```go
message := "brand is required"
res := response.BadRequest(message)
```
Output
```json
{
    "message": "bad request",
    "code": 400,
    "errors": "brand is required"
}
```

### Invalid Input response
```go
message := "brand is required"
myStruct := MyStruct{"struct example", 123, true}
res := response.InvalidInput(message, myStruct)

myStructRes := res.GetData().(MyStruct)
```

### Custom headers
```go
car := &Car{"Ferrari"}
h := response.NewHeaders().Add("Access-Control-Allow-Origin", "*")
res := response.Success("success", car, nil, h)
```
Output
```json
{
    "message": "success",
    "code": 200,
    "data" {
        "brand": "Ferrari"
    }
    "meta": null
}
```