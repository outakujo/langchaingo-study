package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var actionTmp = "Action:%s\nAction Input:%s\n"

var finalAnswerActionTmp = "Final Answer: %s\n"

func main() {
	args := os.Args
	input := strings.Join(args[1:], "")
	r := regexp.MustCompile(`usetool\nObservation:\s*(.+)`)
	matches := r.FindStringSubmatch(input)
	if len(matches) != 0 {
		sps := strings.Split(matches[1], "\\n")
		fmt.Printf(finalAnswerActionTmp, sps[0])
		return
	}
	if strings.Contains(input, "usetool") {
		fmt.Printf(finalAnswerActionTmp, "tool")
	} else if strings.Contains(input, "mytool") {
		fmt.Printf(actionTmp, "mytool", "usetool")
	} else {
		fmt.Printf(finalAnswerActionTmp, "hello")
	}
}
