package switcher

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type SinkInput struct {
	Id     int
	Name   string
	SinkId int
}

func (self *SinkInput) Print() {
	fmt.Printf("%d: %s (%d)\n", self.Id, self.Name, self.SinkId)
}

type PactlParser struct {
}

func NewParser() *PactlParser {
	return &PactlParser{}
}

func (self *PactlParser) Parse(reader io.Reader) (sinks []SinkInput, err error) {
	scanner := bufio.NewScanner(reader)

	var id int = -1
	var sinkId int = -1
	var name string = ""

	for scanner.Scan() {
		line := scanner.Text()

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

		if id != -1 {
			if sinkId == -1 {
				sinkId = self.parseSinkId(line)
			}
			if name == "" {
				name = self.parseName(line)
			}

		}
		if input := self.createInputIfPossible(id, name, sinkId); input != nil {
			sinks = append(sinks, *input)
			id = -1
			name = ""
			sinkId = -1
		}
	}
	if err != nil {
		return
	}
	err = scanner.Err()
	return
}

func (self *PactlParser) createInputIfPossible(id int, name string, sinkId int) *SinkInput {

	if id != -1 && name != "" && sinkId != -1 {
		return &SinkInput{
			Id:     id,
			Name:   name,
			SinkId: sinkId,
		}
	}
	return nil
}

func (self *PactlParser) parseName(line string) (name string) {
	if !strings.Contains(line, "application.process.binary") {
		return
	}

	trimmedLine := strings.TrimSpace(line)
	nameSplit := strings.Split(trimmedLine, "=")

	if len(nameSplit) == 2 {
		name = nameSplit[1]
	}
	return
}

func (self *PactlParser) parseSinkId(line string) int {
	prefix := "Sink:"
	trimmedLine := strings.TrimSpace(line)
	if !strings.HasPrefix(trimmedLine, prefix) {
		return -1
	}
	sinkIdString := strings.TrimPrefix(trimmedLine, prefix)

	sinkId, err := strconv.Atoi(strings.TrimSpace(sinkIdString))
	if err == nil {
		return sinkId
	}
	return -1
}
