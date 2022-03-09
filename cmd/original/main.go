package main

import (
	"fmt"

	"github.com/amobe/go-practice-stacktrace/pkg/handler"
	"github.com/amobe/go-practice-stacktrace/pkg/service"
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
}

func mainWithError() error {
	s := service.NewService()
	h := handler.NewHandler(s)
	res, err := h.HelloOriginal()
	if err != nil {
		return fmt.Errorf("handle hello failed: %w", err)
	}
	fmt.Println(res)
	return nil
}
