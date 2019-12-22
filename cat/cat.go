package main

import (
	"os"

	student ".."
)

func main() {
	args := os.Args[1:]
	if student.StringArrLen(args) == 0 {
		const BUF_SIZE = 512
		in := os.Stdin
		out := os.Stdout
		bufOut := make([]byte, BUF_SIZE)
		n := 0

		for {
			if _, err := in.Read(bufOut[n : n+1]); err != nil {
				break
			}

			if bufOut[n] == '\n' || n == BUF_SIZE {
				if _, err := out.Write(bufOut[0 : n+1]); err != nil {
					break
				}

				n = 0
			} else {
				n++
			}
		}
	}
}