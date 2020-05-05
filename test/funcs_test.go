package test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/11grossmane/mshl"
	. "github.com/smartystreets/goconvey/convey"
)

type S struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func TestFuncs(t *testing.T) {
	Convey("Unmarshaler works as expected", t, func() {
		// normal case
		js := []byte(`{"name":"jo","age":12}`)
		expected := S{}
		err := json.Unmarshal(js, &expected)
		So(err, ShouldBeNil)
		actual, ok := mshl.Unmarshal(js, S{}).(S)
		So(ok, ShouldBeTrue)
		So(actual, ShouldResemble, expected)

		//error case
		js = []byte(`{"word":40,"age":[10]}`)
		expected = S{}
		So(err, ShouldBeNil)
		actual, ok = mshl.Unmarshal(js, S{}).(S)
		So(ok, ShouldBeFalse)
		So(actual, ShouldResemble, expected)
	})
	Convey("Decoder works as expected", t, func() {
		reader := bytes.NewReader([]byte(`{"name":"jo","age":12}`))
		expected := S{}
		err := json.NewDecoder(reader).Decode(&expected)
		So(err, ShouldBeNil)
		reader = bytes.NewReader([]byte(`{"name":"jo","age":12}`))
		actual, ok := mshl.Decode(reader, S{}).(S)
		So(ok, ShouldBeTrue)
		So(actual, ShouldResemble, expected)
	})

	Convey("Marshal works as expected", t, func() {
		expected := []byte(`{"name":"jo","age":12}`)
		inputStruct := S{Name: "jo", Age: 12}
		actual, err := mshl.Marshal(inputStruct)
		So(err, ShouldBeNil)
		So(string(actual), ShouldEqual, string(expected))
	})
}
