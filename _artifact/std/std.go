package std

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Leia() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return line[:len(line)-1]
}

func Leia__int() int64 {
	fval, err := strconv.ParseInt(Leia(), 10, 64)

	if err != nil {
		panic(err)
	}

	return fval
}

func Leia__float() float64 {
	fval, err := strconv.ParseFloat(Leia(), 64)

	if err != nil {
		panic(err)
	}

	return fval
}

func Leia__string() string {
	return Leia()
}

func Escreva(a ...any) {
	fmt.Println(a...)
}
