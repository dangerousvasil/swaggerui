package swaggerui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate go run generate.go

//go:embed embed
var Swagfs embed.FS

func ByteHandler(b []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write(b)
	}
}

// Handler returns a handler that will serve a self-hosted Swagger UI with your spec embedded
func Handler(spec []byte) http.Handler {
	// render the index template with the proper spec name inserted
	static, _ := fs.Sub(Swagfs, "embed")
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger_spec", ByteHandler(spec))
	mux.Handle("/", http.FileServer(http.FS(static)))
	return mux
}
