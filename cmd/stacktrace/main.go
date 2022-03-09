package main

import (
	"fmt"

	"github.com/amobe/go-practice-stacktrace/pkg/handler"
	"github.com/amobe/go-practice-stacktrace/pkg/service"
	"github.com/romanyx/stack"
)

func main() {
	if err := mainWithError(); err != nil {
		printStacktrace(err)
	}
}

func printStacktrace(err error) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %v\n\n", err)
	for _, frame := range stack.Trace(err) {
		fmt.Printf("%s:%d %s()\n", frame.File, frame.Line, frame.Function)
	}
}

func mainWithError() error {
	s := service.NewService()
	h := handler.NewHandler(s)
	res, err := h.HelloStacktrace()
	if err != nil {
		return stack.Errorf("handle hello failed: %w", err)
	}
	fmt.Println(res)
	return nil
}
