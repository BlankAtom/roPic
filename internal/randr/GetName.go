package randr

import (
	"fmt"
	"time"
)

// RenamePicture 根据调用的时间生成一个时间戳并返回
func RenamePicture() string {
	n := time.Now()
	y := n.Year()
	m := n.Month()
	d := n.Day()
	h := n.Hour()
	mm := n.Minute()
	s := n.Second()
	ss := n.Nanosecond()
	str := fmt.Sprintf("%d%02d%02d%02d%02d%02d%03d", y, m, d, h, mm, s, ss/1000)
	return str
}
