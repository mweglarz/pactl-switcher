package switcher

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type PactlParser struct {
}

func NewParser() *PactlParser {
	return &PactlParser{}
}

func (self *PactlParser) Parse(reader io.Reader) (sinks []SinkInput, err error) {

	// bufferedReader := bufio.NewReader(reader)

	scanner := bufio.NewScanner(reader)

	var id int = -1

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line", line)

		if strings.Contains(line, "Sink Input") {
			idSplit := strings.Split(line, "#")

			if len(idSplit) == 2 {
				tempId, err := strconv.Atoi(idSplit[1])
				if err != nil {
					continue
				}
				id = tempId
			}
		}

		if id != -1 && strings.Contains(line, "application.process.binary") {
			trimmedLine := strings.TrimSpace(line)
			nameSplit := strings.Split(trimmedLine, "=")

			if len(nameSplit) == 2 {
				name := nameSplit[1]

				sinks = append(sinks, SinkInput{
					Id:   id,
					Name: name,
				})
				id = -1
			}
		}
	}
	if err != nil {
		return
	}
	err = scanner.Err()
	return
}
