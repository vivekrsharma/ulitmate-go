package chapter3

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "世界 means world"
	var buf [utf8.UTFMax]byte
	fmt.Println(str)
	for i, r := range str {
		rl := utf8.RuneLen(r)
		//fmt.Printf("%d %d", i, rl)
		si := i + rl
		copy(buf[:], str[i:si])
		fmt.Printf("%2d: %q; codepoint %#6x; encoded bytes: %#v \n",i, r, r, buf[:rl])

	}
}
