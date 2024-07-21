package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	numTimes   int
	printUsage bool
}

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|-help]

A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])

// 프로그램의 표준 출력을 담당하는 함수
func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

// 입력 매개변수로 문자열의 슬라이스를 받고 config타입과 error타입 2개를 반환하는 함수
func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "-help" {
		c.printUsage = true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes

	return c, nil
}

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

// 입력값을 검사하는 함수
func validateArgs(c config) error {
	if c.numTimes <= 0 && !c.printUsage {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

// 사용자의 입력에 따라 프로그램 사용법을 출력하거나 이름을 입력받아 인사 메시지를 지정된 횟수만큼 출력하는 함수
func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

// 입력받은 이름을 사용하여 인사 메시지를 생성하고, config 구조체에 정의된 횟수만큼 메시지를 출력하는 함수
func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

// 프로그램의 main함수
func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
