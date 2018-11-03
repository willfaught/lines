package lines

import (
	"bufio"
	"io"
)

func Lines(in io.Reader) (<-chan []byte, error) {
	b := make(chan []byte)
	r := bufio.NewScanner(in)
	r.Split(bufio.ScanLines)
	for r.Scan() {
		r.Bytes()
	}
}

func Unlines(out io.Writer, c <-chan []byte) error {

}
