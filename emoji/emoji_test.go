package emoji

import (
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	s := "a" + "🏴󠁧󠁢󠁥󠁮󠁧󠁿"
	s += "是"
	b := Match("❤a❤")
	fmt.Println("match:", b)
	fmt.Println("len:", Length("爱❤️"))
}
