package something

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOtherExample(t *testing.T) {
	Convey("What about another example?", t, func() {
		Convey("that does something else...", func() {
			So(1, ShouldEqual, 1)
		})
	})
}

// func TestRegular(t *testing.T) {
// 	s := []byte{}[40]
// 	fmt.Println(s)
// }

func TestExample(t *testing.T) {
	Convey("Subject: Integer incrementation and decrementation", t, func() {
		x := 0

		Convey("Given a starting integer value", func() {
			x = 42

			Convey("When incremented", func() {
				x++

				Convey("The value should be greater by one", func() {
					So(x, ShouldEqual, 43)
					So(x, ShouldEqual, -1234)
				})
				Convey("The value should NOT be what it used to be", func() {
					// So(x, ShouldNotEqual, 42)
				})
			})
			Convey("When decremented", func() {
				x--

				Convey("The value should be lesser by one", func() {
					So(x, ShouldEqual, 41)
				})
				Convey("The value should NOT be what it used to be", func() {
					s := []byte{}[40]
					So(s, ShouldEqual, 1234)
					// So(x, ShouldNotEqual, 42)
				})
			})
			Reset(func() {
				x = 0
			})
		})
	})
}

/*
Output:

### file/to/test.go - Line 42 - TestExample ###

	- Subject: Integer incrementation and decrementation
		- Given a starting integer value
			- When incremented
				- The value should be greater by one
				- The value should NOT be what it used to be
			- When decremented
				- The value should be lesser by one


## Running Total: ##

1 Stor[y|ies] with 4 Assertion[s] (1 Failed)[, 0 Errors]

*/