package base62

import (
	"fmt"
	assert "github.com/pilu/miniassert"
	"math"
	"testing"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, "0", Encode(0))
	assert.Equal(t, "1B", Encode(99))
	assert.Equal(t, "lYGhA16ahyf", Encode(math.MaxUint64))
}

func TestDecode(t *testing.T) {
	assert.Equal(t, 0, int(Decode("0")))
	assert.Equal(t, 99, int(Decode("1B")))
	assert.Equal(t, int64(-1), int64(Decode("lYGhA16ahyf")))
}

func ExampleEncode() {
	fmt.Println(Encode(99))
	// Output:
	// 1B
}

func ExampleDecode() {
	fmt.Println(Decode("1B"))
	// Output:
	// 99
}
