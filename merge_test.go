package envlib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMerge(t *testing.T) {

	Convey("With the given environment", t, func() {

		nameset := NewNameset("ENVLIB_TEST1", "ENVLIB_TEST2")
		environ := []string{"ENVLIB_TEST1=value1", "ENVLIB_TEST2=value2", "ENVLIB_TEST3=value3"}

		Init(environ)
		Select(nameset)

		Convey("When merging with a environment file", func() {
			MergeWith("testdata/etc_environment_example")

			Convey("Data from file is loaded", func() {
				So(Getenv("ENVLIB_TESTDATA1"), ShouldEqual, "loaded")
			})

			Convey("And the selected environment have preference over the file", func() {
				So(Getenv("ENVLIB_TEST1"), ShouldEqual, "value1")
			})

		})

	})

}
