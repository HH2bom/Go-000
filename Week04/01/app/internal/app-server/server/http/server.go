package http

import (
	"context"
	"net/http"

	libraryapi "app/internal/app-server/api/library/v1"
	libraryserv "app/internal/app-server/service/library/v1"
	"app/internal/pkg/appcontext"
	"app/internal/pkg/conf"
)

func Server(ctx context.Context, cfg conf.C) *http.Server {

	// 依赖待使用 wire
	appCtx := appcontext.New(appcontext.Config(cfg))

	libraryServ := libraryserv.NewLibraryService(*appCtx)

	mux := http.NewServeMux()
	mux.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		libraryServ.GetBook(r.Context(), libraryapi.GetBookRequest{})
	})

	return &http.Server{Addr: ":80", Handler: mux}
}
