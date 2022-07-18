package main

import (
	"strings"

	"github.com/jduhamel/readline"
)

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 "> ",
		HistoryFile:            "/tmp/readline-multiline",
		DisableAutoSaveHistory: false,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	var cmds []string
	index := 0
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cmds = append(cmds, line)
		if !strings.HasSuffix(line, ";") {
			rl.SetPrompt(">>> ")
			index += 1
			continue
		}

		cmd := strings.Join(cmds, " ")
		cmds = cmds[:0]
		rl.SetPrompt("> ")

		if index != 0 {
			rl.SaveReplaceHistory(cmd, index)
		}
		index = 0

	}
}
