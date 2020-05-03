package test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/11grossmane/mshl/pkg"
	. "github.com/smartystreets/goconvey/convey"
)

type S struct {
	name string `json:"name,omitempty"`
	age  int    `json:"age,omitempty"`
}

func TestFuncs(t *testing.T) {
	Convey("Unmarshaler works as expected", t, func() {
		js := []byte(`{"name":"jo","age":12}`)
		expected := S{}
		err := json.Unmarshal(js, &expected)
		So(err, ShouldBeNil)

		actual := pkg.Unmarshal(js, S{}).(S)
		So(actual, ShouldResemble, expected)
	})
	Convey("Decoder works as expected", t, func() {
		reader := bytes.NewReader([]byte(`{"name":"jo","age":12}`))
		expected := S{}
		err := json.NewDecoder(reader).Decode(&expected)
		So(err, ShouldBeNil)

		actual := pkg.Decode(reader, S{}).(S)
		So(actual, ShouldResemble, expected)
	})
}
