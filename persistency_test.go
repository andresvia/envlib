package envlib

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestPersistency(t *testing.T) {

	Convey("Provided a filename", t, func() {
		file := "testdata/generated/persistency"
		Convey("The file is removed", func() {
			So(os.RemoveAll(file), ShouldBeNil)
		})
		Convey("Save the environment to the file", func() {
			Init([]string{"var1=val1", "var2=val2"})
			So(Persist(file), ShouldBeNil)
		})
		Convey("Recover the environment from the file", func() {
			Init([]string{})
			MergeWith(file)
			So(Getenv("var1"), ShouldEqual, "val1")
			So(Getenv("var2"), ShouldEqual, "val2")
		})

	})

}
