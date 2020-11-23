package array_test

import (
	"fmt"

	"github.com/openacid/slim/array"
)

func Example() {

	// arr[0]   = 12
	// arr[5]   = 15
	// arr[9]   = 19

	indexes := []int32{0, 5, 9}
	elts := []uint32{12, 15, 19}

	arr, err := array.New(indexes, elts)
	if err != nil {
		fmt.Printf("Init compacted array error:%s\n", err)
		return
	}

	val, found := arr.Get(5)
	fmt.Printf("get arr[%d] value:%v found: %t\n", 5, val, found)

	// Output:
	// get arr[5] value:15 found: true
}
