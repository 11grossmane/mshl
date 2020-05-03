package mshl

import (
	"bytes"

	"github.com/11grossmane/mshl"
)

type S struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func ExampleMarshal() {
	marshaledJSON, err := mshl.Marshal(S{Name: "jo", age: 12})
}

func ExampleUnmarshal() {
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	someBytes := []byte(`{"name":"jo","age":12}`)

	myStruct, err := mshl.Unmarshal(someBytes, S{})

}

func ExampleDecode() {
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	someReader := bytes.NewReader([]byte(`{"name":"jo","age":12}`))

	myStruct, err := mshl.Decode(someReader, S{})

}
