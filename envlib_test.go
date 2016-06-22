package envlib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestKnockOff(t *testing.T) {

	Convey("Given a envlib variable passed with empty value", t, func() {

		nameset := NewNameset("ENVLIB_TEST1", "ENVLIB_TESTDATA1")
		environ := []string{"ENVLIB_TESTDATA1="}
		result := "ENVLIB_TEST1=overriden"
		Init(environ)

		Convey("It will knock-off existing variables after merging", func() {
			MergeWith("testdata/etc_environment_example")
			Select(nameset)
			So(ToString(), ShouldEqual, result)
		})

	})
}

func TestInit(t *testing.T) {
	Convey("Given a broken environ", t, func() {
		environ := []string{"broken_environ"}
		Convey("Init shall panic", func() {
			So(func() { Init(environ) }, ShouldPanic)
		})
	})
}
