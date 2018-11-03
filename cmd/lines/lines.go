package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func drop(r io.Reader, w io.Writer, n int) error {
	s := bufio.NewScanner(r)
	bw := bufio.NewWriter(w)
	lines := make([][]byte, 0, n)
	var i int
	for s.Scan() {
		line := s.Bytes()
		if len(lines) < n {
			lines = append(lines, line)
			i = (i + 1) % n
		} else {
			old := lines[i]
			lines[i] = line
			if _, err := bw.Write(old); err != nil {
				return err
			}
			if _, err := fmt.Fprintln(bw); err != nil {
				return err
			}
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	return bw.Flush()
}

func run() error {
	l := len(os.Args)
	if l < 2 {
		return errors.New("command is missing")
	}
	switch cmd := os.Args[1]; cmd {
	case "drop":
		if l < 3 {
			return errors.New("count is missing")
		}
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return err
		}
		if n < 0 {
			return fmt.Errorf("%d is invalid", n)
		}
		return drop(os.Stdin, os.Stdout, n)
	default:
		return fmt.Errorf("%s command invalid", cmd)
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": error: ")
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
