
import "strings"

// convention use sufix er  -talker = anything that talks
var talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	// strings.repeat (quantas vezes vai repetir uma string)
	return strings.Repeat("pew", int(1))
}
