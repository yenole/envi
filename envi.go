package main

import (
	"fmt"
	"os"
)

func parseCommand(args []string) (string, error) {
	if cmd, ok := cmds[args[0]]; ok {
		if cmd.Args <= uint32(len(args)) {
			return cmd.Handle.handle(args[1:])
		}
		return "", fmt.Errorf("Number of parameters does not match")
	}
	return "", fmt.Errorf("Unknown command")
}

func main() {
	if len(os.Args) > 1 {
		out, err := parseCommand(os.Args[1:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
		return
	}
	fmt.Println(help())
}
