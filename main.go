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

func pwd(args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}

func echo(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func cat(args []string) error {
	if len(args) < 1 {
		return errors.New("filename required")
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func mkdir(args []string) error {
	if len(args) < 1 {
		return errors.New("directory name required")
	}
	return os.Mkdir(args[0], 0755)
}

func help(args []string) error {
	fmt.Println("Available сommands:")
	fmt.Println("cd     - change directory")
	fmt.Println("ls     - list directory contents")
	fmt.Println("pwd    - print current working directory")
	fmt.Println("echo   - print arguments")
	fmt.Println("cat    - display file content")
	fmt.Println("mkdir  - create a new directory")
	fmt.Println("help   - display this help message")
	fmt.Println("exit   - exit the shell")
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
	case "pwd":
		return pwd(args[1:])
	case "echo":
		return echo(args[1:])
	case "cat":
		return cat(args[1:])
	case "mkdir":
		return mkdir(args[1:])
	case "help":
		return help(args[1:])
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
