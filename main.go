package main

import (
	"GoFxvsNon/server"
	"context"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(http.NewServeMux),
		fx.Invoke(server.New),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(
	lifecycle fx.Lifecycle, mux *http.ServeMux,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go http.ListenAndServe(":8080", mux)
				return nil
			},
		},
	)
}
