package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

// Unmarshal allows you to unmarshal json with one line of code
func Unmarshal(js []byte, empty interface{}) interface{} {
	//grabbing type
	thingType := reflect.TypeOf(empty)

	//creating new empty pointer from type
	newThing := reflect.New(thingType).Interface()

	//unmarshaling from json
	json.Unmarshal(js, &newThing)

	//grabbing value from pointer, must use reflect.indirect because it's newThing is interface{}
	//then turning reflect.Value back to interface{}
	indirectedValue := reflect.Indirect(reflect.ValueOf(newThing)).Interface()
	return indirectedValue
}

// Marshal marshales struct into json, same as json.Marshal
func Marshal(any interface{}) ([]byte, error) {
	js, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}
	return js, nil
}

// Decode allows you to decode json with one line of code
func Decode(reader io.Reader, empty interface{}) interface{} {
	//grabbing type
	thingType := reflect.TypeOf(empty)

	//creating new empty pointer from type
	newThing := reflect.New(thingType).Interface()

	//unmarshaling from json
	json.NewDecoder(reader).Decode(&newThing)

	//grabbing value from pointer, must use reflect.indirect because it's newThing is interface{}
	//then turning reflect.Value back to interface{}
	indirectedValue := reflect.Indirect(reflect.ValueOf(newThing)).Interface()
	fmt.Println("indirected", indirectedValue)
	return indirectedValue
}
