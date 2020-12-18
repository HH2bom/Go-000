package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"

	"app/internal/pkg/conf"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"app/internal/app-server/server/http"
	"app/internal/app-server/server/rpc"
)

var (
	isRunHTTPServer, isRunRPCServer bool
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "app server",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("server called")

		cfg, err := initConfig(cfgFile)
		if err != nil {
			return err
		}

		ctx := context.Background()
		defer ctx.Done()

		// 信号监听
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)

		var g *errgroup.Group

		switch {
		case isRunHTTPServer:
			g = runHTTPServer(ctx, *cfg, sig)
		case isRunRPCServer:
			g = runRPCServer(ctx, *cfg, sig)
		default:
			return errors.New("server type not found")
		}

		return g.Wait()
	},
}

func runHTTPServer(ctx context.Context, cfg conf.C, sig <-chan os.Signal) *errgroup.Group {
	g, gCtx := errgroup.WithContext(ctx)
	server := http.Server(ctx, cfg)
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

	return g
}

func runRPCServer(ctx context.Context, cfg conf.C, sig <-chan os.Signal) *errgroup.Group {
	g, gCtx := errgroup.WithContext(ctx)
	_, run, stop := rpc.Server(ctx, cfg)
	g.Go(run)
	g.Go(func() error {
		select {
		case <-gCtx.Done(): // maybe server down
			fmt.Printf("ctx done by: %v\n", gCtx.Err())
			return gCtx.Err()
		case value := <-sig:
			fmt.Printf("catch sig: %v\n", value)
			stop()
			return nil
		}
	})

	return g
}

func init() {
	rootCmd.AddCommand(serverCmd)
	sf := serverCmd.Flags()

	sf.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.app.yaml)")
	sf.BoolVar(&isRunHTTPServer, "http", false, "run http server")
	sf.BoolVar(&isRunRPCServer, "rpc", false, "run rpc server")
}
