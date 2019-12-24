package main

import (
	"fmt"
	"io/ioutil"
	"os"

	student ".."
)

func main() {
	args := os.Args[1:]
	if student.StringArrLen(args) == 0 {
		ioutil.ReadAll(os.Stdin)
		defer os.Stdin.Close()
	}

	if contains(args, "-c") {
		if student.StringArrLen(args) == 3 {
			nbr, tailErr := Atoi(args[1])
			if tailErr != nil {
				student.PrintString(tailErr.Error())
				os.Exit(tailErr.Code())
			}

			file, err := os.OpenFile(args[2], os.O_RDONLY, 0700)
			if err != nil {
				student.PrintString(err.Error())
				os.Exit(5)
			}
			defer file.Close()

			switch args[1][0] {
			case '+':
				//file.Seek(int64(nbr), 0)

				stat, _ := file.Stat()

				var bytes []byte
				var i int64 = 0
				for ; i <= stat.Size()-int64(nbr); i++ {
					bytes = append(bytes, '\x00')
				}

				if _, err := file.Read(bytes); err != nil {
					student.PrintString(err.Error())
					os.Exit(5)
				}

				if _, err := fmt.Printf("%s", bytes); err != nil {
					student.PrintString(err.Error())
					os.Exit(5)
				}
			case '-':
				//file.Seek(int64(nbr), 2)

				stat, _ := file.Stat()

				var bytes []byte
				var i int64 = 0
				for ; i <= int64(nbr); i++ {
					bytes = append(bytes, '\x00')
				}

				if _, err := file.ReadAt(bytes, stat.Size()-int64(nbr)); err != nil {
					student.PrintString(err.Error())
					os.Exit(5)
				}

				if _, err := fmt.Printf("%s", bytes); err != nil {
					student.PrintString(err.Error())
					os.Exit(5)
				}
			default:
				//file.Seek(int64(nbr), 2)

				stat, _ := file.Stat()

				var bytes []byte
				var i int64 = 0
				for ; i <= int64(nbr); i++ {
					bytes = append(bytes, '\x00')
				}
				if _, err := file.ReadAt(bytes, stat.Size()-int64(nbr)); err != nil {
					student.PrintString(err.Error())
					os.Exit(5)
				}

				if _, err := fmt.Printf("%s", bytes); err != nil {
					student.PrintString(err.Error())
					os.Exit(5)
				}
			}
		} else {
			student.PrintString("ztail: option requires an argument -- 'c'")
			os.Exit(22)
		}
	}
}

func contains(arr []string, s string) bool {
	for _, str := range arr {
		if str == s {
			return true
		}
	}

	return false
}

func Atoi(s string) (int, *tailError) {
	result := 0

	str := []rune(s)
	isNegative := false
	if s[0] == '-' || s[0] == '+' {
		str = []rune(s[1:])
		if s[0] == '-' {
			isNegative = true
		}
	}
	len := student.StringLen(str)

	power := 1

	for i := len - 1; i >= 0; i-- {
		//z01.PrintRune(str[i])
		if str[i] >= '0' && str[i] <= '9' {
			result += int(str[i]-48) * power
			power *= 10
		} else {
			return 0, newError("ztail: invalid number of bytes: ‘"+s+"’", 22)
		}
	}

	if isNegative {
		result *= -1
	}

	return result, nil
}

func newError(text string, code int) *tailError {
	return &tailError{text, code}
}

type tailError struct {
	s    string
	code int
}

func (e tailError) Error() string {
	return e.s
}

func (e tailError) Code() int {
	return e.code
}
