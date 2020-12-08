package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	g := &errgroup.Group{}

	g.Go(func() error {
		return errors.New("first error")
	})
	g.Go(func() error {
		return errors.New("second error")
	})

	// 注意，只能收到第一个 error
	if err := g.Wait(); err != nil {
		fmt.Printf("catch error: %v", err)
	}
}
