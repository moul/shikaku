package shikaku

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleShikakuMap_String() {
	rand.Seed(42)
	shikakuMap := NewShikakuMap(8, 4, 0, 0)
	if err := shikakuMap.GenerateBlocks(5); err != nil {
		panic(err)
	}
	fmt.Println(shikakuMap.String())
	// Output:
	// T 8 4
	// 8 1 1
	// 4 3 1
	// 8 4 2
	// 4 3 2
	// 8 4 1
}

func ExampleShikakuMap_DrawMap() {
	rand.Seed(42)
	shikakuMap := NewShikakuMap(8, 4, 0, 0)
	if err := shikakuMap.GenerateBlocks(5); err != nil {
		panic(err)
	}
	fmt.Println(shikakuMap.DrawMap())
	// Output:
	// +---+---+---+---+---+---+---+---+
	// |                               |
	// +   +   +   +   +   +   +   +   +
	// |     8       4   8             |
	// +   +   +   +   +   +   +   +   +
	// |             4   8             |
	// +   +   +   +   +   +   +   +   +
	// |                               |
	// +---+---+---+---+---+---+---+---+
}

func ExampleShikakuMap_DrawSolution() {
	rand.Seed(42)
	shikakuMap := NewShikakuMap(8, 4, 0, 0)
	if err := shikakuMap.GenerateBlocks(5); err != nil {
		panic(err)
	}
	fmt.Println(shikakuMap.DrawSolution())
	// Output:
	// +---+---+---+---+---+---+---+---+
	// |       |       |               |
	// +   +   +   +   +   +   +   +   +
	// |       |       |               |
	// +   +   +---+---+---+---+---+---+
	// |       |       |               |
	// +   +   +   +   +   +   +   +   +
	// |       |       |               |
	// +---+---+---+---+---+---+---+---+
}

func ExampleShikakuMap_DrawEmptyAsciiMap() {
	rand.Seed(42)
	shikakuMap := NewShikakuMap(8, 4, 0, 0)
	if err := shikakuMap.GenerateBlocks(5); err != nil {
		panic(err)
	}
	fmt.Println(strings.Join(shikakuMap.DrawEmptyAsciiMap(), "\n"))
	// Output:
	// +---+---+---+---+---+---+---+---+
	// +   |   |   |   |   |   |   |   |
	// +---+---+---+---+---+---+---+---+
	// +   |   |   |   |   |   |   |   |
	// +---+---+---+---+---+---+---+---+
	// +   |   |   |   |   |   |   |   |
	// +---+---+---+---+---+---+---+---+
	// +   |   |   |   |   |   |   |   |
	// +---+---+---+---+---+---+---+---+
}

func Test(t *testing.T) {
	Convey("Testing package", t, func() {
		rand.Seed(42)
		shikakuMap := NewShikakuMap(8, 4, 0, 0)
		err := shikakuMap.GenerateBlocks(5)
		So(err, ShouldBeNil)
		So(shikakuMap.String(), ShouldNotBeEmpty)
		So(shikakuMap.DrawMap(), ShouldNotBeEmpty)
		So(shikakuMap.DrawSolution(), ShouldNotBeEmpty)

		rand.Seed(42)
		shikakuMap2 := NewShikakuMap(8, 4, 0, 0)
		err = shikakuMap2.GenerateBlocks(5)
		So(err, ShouldBeNil)
		So(shikakuMap.String(), ShouldEqual, shikakuMap2.String())
		So(shikakuMap.DrawMap(), ShouldEqual, shikakuMap2.DrawMap())
		So(shikakuMap.DrawSolution(), ShouldEqual, shikakuMap2.DrawSolution())
	})
}
