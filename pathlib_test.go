package pathlib

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	p := New("../")

	fmt.Println(p.Absolute())
	fmt.Println(p.Cwd())
	fmt.Println(p.Exists())
	fmt.Println(p.Parent())
}
