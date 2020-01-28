package main


const size = 1024

func main() {
	s:= "HELLO"
	var a [size]int
	stackCopy(&s, 0 , a)
}

func stackCopy(s *string , c int , a [size]int) {
	println(c, s, *s)
	c++
	if c == 10 {
		return
	}
	stackCopy(s, c, a)
}
