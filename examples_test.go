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

	marshaledJSON:= mshl.Marshal(S{Name: "jo", Age: 12})
	fmt.Println(string(marshaledJSON))
}

func ExampleUnmarshal() {
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	someBytes := []byte(`{"name":"jo","age":12}`)

	myStruct, ok := mshl.Unmarshal(someBytes, S{}).(S)
	if ok != true {
		log.Panic("Could not assert type")
	}
	fmt.Println(myStruct)

}

func ExampleDecode() {
	type S struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	someReader := bytes.NewReader([]byte(`{"name":"jo","age":12}`))

	myStruct, ok := mshl.Decode(someReader, S{}).(S)
	if ok != true {
		log.Panic("Could not assert type")
	}
	fmt.Println(myStruct)

}
