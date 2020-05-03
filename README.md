# Go Package For Quick JSON Unmarshaling
- Go package to Marshal, Unmarshal, and Decode JSON with one line of code

## Install
`go get -u github.com/11grossmane/mshl`

## Basic Usage

### Unmarshal
```go
type S struct {
  Name string `json:"name,omitempty"`
  Age  int    `json:"age,omitempty"`
	}

someBytes := []byte(`{"name":"jo","age":12}`)
myStruct, err := mshl.Unmarshal(someBytes, S{})
```

### Decode Reader
```go
type S struct {
  Name string `json:"name,omitempty"`
  Age  int    `json:"age,omitempty"`
}

req := &http.Request{}
myStruct, err := mshl.Decode(req.Body, S{})
```

### Marshal
this is exactly the same as json.Marshal
```go
type S struct {
  Name string `json:"name,omitempty"`
  Age  int    `json:"age,omitempty"`
}

marshaledJSON, err := mshl.Marshal(S{Name: "jo", age: 12})
```


## License
MIT