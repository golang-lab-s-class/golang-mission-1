package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

// 사용자의 이름을 입력받는 함수
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprint(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}
	return name, nil
}
