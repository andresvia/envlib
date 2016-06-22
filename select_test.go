package envlib

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSelect(t *testing.T) {

	Convey("Given an environ Select shall copy the listed environment variables", t, func() {
		nameset := NewNameset("ENVLIB_TEST1", "ENVLIB_TEST2")
		environ := []string{"ENVLIB_TEST1=value1", "ENVLIB_TEST2=value2", "ENVLIB_TEST3=value3"}
		result := []string{"ENVLIB_TEST1=value1", "ENVLIB_TEST2=value2"}
		Init(environ)
		Select(nameset)
		value1 := Getenv("ENVLIB_TEST1")
		value2 := Getenv("ENVLIB_TEST2")
		value3 := Getenv("ENVLIB_TEST3")
		So(value1, ShouldEqual, "value1")
		So(value2, ShouldEqual, "value2")
		So(value3, ShouldEqual, "")
		So(fmt.Sprintf("%s", result), ShouldEqual, fmt.Sprintf("%s", Environ()))
	})

	Convey("Given an environ with duplicated entries, the last one wins", t, func() {
		nameset := NewNameset("ENVLIB_TEST")
		environ := []string{"ENVLIB_TEST=first_value", "ENVLIB_TEST=second_value"}
		result := []string{"ENVLIB_TEST=second_value"}
		Init(environ)
		Select(nameset)
		So(Getenv("ENVLIB_TEST"), ShouldEqual, "second_value")
		So(fmt.Sprintf("%s", result), ShouldEqual, fmt.Sprintf("%s", Environ()))
	})

	Convey("Given an environ with missing entries from the Nameset, they should be ignored", t, func() {
		nameset := NewNameset("ENVLIB_TEST", "ENVLIB_MISSING")
		environ := []string{"ENVLIB_TEST=value"}
		result := []string{"ENVLIB_TEST=value"}
		Init(environ)
		Select(nameset)
		So(fmt.Sprintf("%s", result), ShouldEqual, fmt.Sprintf("%s", Environ()))
	})

}
