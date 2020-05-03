package mshl

import (
	"bytes"
	"fmt"
	"log"

	"github.com/11grossmane/mshl"
)

func ExampleMarshal() {
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	marshaledJSON, err := mshl.Marshal(S{Name: "jo", Age: 12})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(marshaledJSON)
}

func ExampleUnmarshal() {
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	someBytes := []byte(`{"name":"jo","age":12}`)

	myStruct, err := mshl.Unmarshal(someBytes, S{})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(myStruct)

}

func ExampleDecode() {
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	someReader := bytes.NewReader([]byte(`{"name":"jo","age":12}`))

	myStruct, err := mshl.Decode(someReader, S{})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(myStruct)

}
