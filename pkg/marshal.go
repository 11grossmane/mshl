package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Marshal(js []byte, empty interface{}) interface{} {
	//grabbing type
	thingType := reflect.TypeOf(empty)
	fmt.Println("thing", thingType)

	//creating new empty pointer from type
	newThing := reflect.New(thingType).Interface()

	//unmarshaling from json
	json.Unmarshal(js, &newThing)

	//grabbing value from pointer, must use reflect.indirect because it's newThing is interface{}
	//then turning reflect.Value back to interface{}
	indirectedValue := reflect.Indirect(reflect.ValueOf(newThing)).Interface()
	fmt.Println("indirected", indirectedValue)
	return indirectedValue
}
