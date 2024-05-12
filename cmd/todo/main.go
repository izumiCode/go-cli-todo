package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/izumiCode/go-cli-todo"
)

const (
	todoFile = "todos.json"
)

func addTodo(todoFile string, todos *todo.Todos) {

	var task, err = getInput(os.Stdin, flag.Args()...)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	todos.Add(task)
	//save the todo.json
	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func completeTodo(complete int, todos *todo.Todos) {
	var err = todos.Complete(complete)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	//save the todo.json
	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func deleteTodo(delete int, todos *todo.Todos) {
	var err = todos.Delete(delete)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	//save the todo.json
	err = todos.Store(todoFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {

	var add = flag.Bool("add", false, "add new todo")
	var complete = flag.Int("c", 0, "mark todo as completed")
	var delete = flag.Int("d", 0, "delete a todo")
	var list = flag.Bool("l", false, "list all todos")
	flag.Parse()

	var todos = &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		addTodo(todoFile, todos)
	case *complete > 0:
		completeTodo(*complete, todos)
	case *delete > 0:
		deleteTodo(*delete, todos)
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)

	}
}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	var scanner = bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}
	var text = scanner.Text()
	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}
	return text, nil
}
