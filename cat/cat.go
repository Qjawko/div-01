package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"

	student ".."
)

func main() {
	args := os.Args[1:]
	if student.StringArrLen(args) == 0 { //esli kol-vo args ravni nulyu
		const BUF_SIZE = 512 //maksimum schitivaemyh simvolov
		in := os.Stdin 
		out := os.Stdout
		bufOut := make([]byte, BUF_SIZE) //sozdaem massiv v kotorom budet nahoditsya nashi vvedenie dannie/hello 
		n := 0 //index simvola

		for { //endless for loop
			if _, err := in.Read(bufOut[n : n+1]); err != nil { //proverka na novyi symbol
				break
			}

			if bufOut[n] == '\n' || n == BUF_SIZE { // if current symbol is \n (enter) then print out an array bufOut
				if _, err := out.Write(bufOut[0 : n+1]); err != nil { // printing
					break
				}

				n = 0
			} else { //n++
				n++
			}
		}
	} else {
		for _, file := range args { //in args nahoditsya nazvaniya filov
			bytes, err := ioutil.ReadFile(file) //file - nazvanie fila kotoroe nado pro4itat
			if err != nil {
				os.Stdout.WriteString(err.Error() + "\n")
				return
			}
			os.Stdout.Write(bytes) //bytes eto info kotoroya v file, vivod
			z01.PrintRune('\n')
		}
	}
}
