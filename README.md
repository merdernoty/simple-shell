# Simple Shell

This is a simple shell terminal written in Go, demonstrating a basic implementation of built-in commands. The project includes the following commands:

- **cd**: Change directory.
- **ls**: List files and directories.
- **pwd**: Print the current working directory.
- **echo**: Print provided arguments.
- **cat**: Display the contents of a file.
- **mkdir**: Create a new directory.
- **help**: Show help information for built-in commands.
- **exit**: Exit the shell.

## Features

- **Built-in Commands:** Custom commands are implemented directly within the shell instead of calling external utilities.
- **Extensibility:** Easily add new commands by modifying the input handling function.
- **Simplicity:** Minimalistic code that demonstrates the core concepts of building a shell in Go.

## Prerequisites

- [Go](https://golang.org/dl/) version 1.16 or higher

## Installation and Build

1. **Clone the repository:**

   ```bash
   git clone https://github.com/merdernoty/simple-shell.git
   cd simple-shell
