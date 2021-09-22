package engine

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func Engine(input string, output string) error {
	f, err := os.OpenFile(input, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}

	fileReader := bufio.NewReader(f)
	content := ""
	var result int
	for {
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}
		if regexp.MustCompile(`^[0-9]+[-+][0-9]+[=?]`).MatchString(string(line)) {
			operands := regexp.MustCompile(`[0-9]+`).FindAllStringSubmatch(string(line), -1)
			operator := regexp.MustCompile(`[+-]`).FindStringSubmatch(string(line))
			n1, _ := strconv.Atoi(operands[0][0])
			n2, _ := strconv.Atoi(operands[1][0])
			switch operator[0] {
			case "-":
				result = n1 - n2
			case "+":
				result = n1 + n2
			}
			content = content + operands[0][0] + operator[0] + operands[1][0] + "=" + strconv.Itoa(result) + "\n"
		}
	}

	err = ioutil.WriteFile(output, []byte(content), 0777)
	if err != nil {
		panic(err)
	}
	return nil
}
