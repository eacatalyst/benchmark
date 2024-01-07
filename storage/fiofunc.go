package benchmark

import (
	"fmt"
	"os"
	"os/exec"
)

//https://www.educative.io/answers/how-to-execute-linux-commands-in-golang

func Execute(cmd string,
	arg01 string,
	arg02 string,
	arg03 string,
	arg04 string,
	arg05 string,
	arg06 string,
	arg07 string,
	arg08 string,
	arg09 string,
	arg10 string,
	arg11 string,

) {
	out, err := exec.Command(cmd,
		arg01,
		arg02,
		arg03,
		arg04,
		arg05,
		arg06,
		arg07,
		arg08,
		arg09,
		arg10,
		arg11).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(output)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
