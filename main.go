package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(":> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("Path required")

func ls(args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "."
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	return nil
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "ls":
		return ls(args[1:])
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
