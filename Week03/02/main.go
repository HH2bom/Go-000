package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 信号监听
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	server := &http.Server{Addr: ":80", Handler: http.NewServeMux()}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(server.ListenAndServe)
	g.Go(func() error {
		select {
		case <-gCtx.Done(): // maybe server down
			fmt.Printf("ctx done by: %v\n", gCtx.Err())
			return gCtx.Err()
		case value := <-sig:
			fmt.Printf("catch sig: %v\n", value)

			timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*2)
			defer cancel()

			err := server.Shutdown(timeoutCtx) // 不等待请求实际结束
			if err != nil {
				fmt.Printf("shutdown server catch err :%v\n", err)
			}
			return err
		}
	})

	fmt.Printf("server down: %+v\n", g.Wait())
}
