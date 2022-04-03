package main

import (
	"fmt"
	"time"
)

func main() {
	//s := make([]byte, 62)
	//l := 0
	//for i := '0'; i <= '9'; i++ {
	//	s[l] = byte(i)
	//	l++
	//}
	//for i := 'a'; i <= 'z'; i++ {
	//	s[l] = byte(i)
	//	l++
	//}
	//for i := 'A'; i <= 'Z'; i++ {
	//	s[l] = byte(i)
	//	l++
	//}
	//
	//for j := l - 1; j >= 1; j-- {
	//	t := randr.Int() % j
	//	s[t], s[j] = s[j], s[t]
	//}
	n := time.Now()
	y := n.Year()
	m := n.Month()
	d := n.Day()
	h := n.Hour()
	mm := n.Minute()
	s := n.Second()
	ss := n.Nanosecond()
	str := fmt.Sprintf("%d%02d%02d%02d%02d%02d%03d", y, m, d, h, mm, s, ss/1000)

	println(str)
}
