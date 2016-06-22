package envlib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestString(t *testing.T) {

	Convey("The environment can be converted to string for saving into files", t, func() {
		nameset := NewNameset("ENVLIB_TEST1", "ENVLIB_TESTDATA1")
		environ := []string{"ENVLIB_TEST1=value1", "ENVLIB_TEST2=value2"}
		result := "ENVLIB_TESTDATA1=loaded\nENVLIB_TEST1=value1"
		Init(environ)
		MergeWith("testdata/etc_environment_example")
		Select(nameset)
		So(ToString(), ShouldEqual, result)
	})

}
