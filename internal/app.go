package internal

import (
	"cmp"
	"embed"
	"fmt"
	"net/http"
	"os"

	"cortico/internal/links"
	"cortico/public"

	"github.com/jmoiron/sqlx"
	"github.com/leapkit/leapkit/core/db"
	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server"
	_ "github.com/lib/pq"
)

var (
	//go:embed **/*.html **/*.html *.html
	tmpls embed.FS

	DatabaseURL = cmp.Or(os.Getenv("DATABASE_URL"), "postgres://postgres:postgres@localhost:5432/cortico?sslmode=disable")
	DriverName  = "postgres"
)

// Server interface exposes the methods
// needed to start the server in the cmd/app package
type Server interface {
	Addr() string
	Handler() http.Handler
}

func New() Server {
	// Creating a new server instance with the
	// default host and port values.
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
		server.WithSession(
			cmp.Or(os.Getenv("SESSION_SECRET"), "d720c059-9664-4980-8169-1158e167ae57"),
			cmp.Or(os.Getenv("SESSION_NAME"), "leapkit_session"),
		),
		server.WithAssets(public.Files),
	)

	// Application services
	if err := AddServices(r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),

		render.WithDefaultLayout("layout.html"),
	))

	r.HandleFunc("GET /{$}", links.Index)
	r.HandleFunc("POST /short", links.ShortURL)
	r.HandleFunc("GET /{short_url}", links.ShortUrlRedirect)

	return r
}

// DB is the database connection builder function
// that will be used by the application based on the driver and
// connection string.
func DB() (*sqlx.DB, error) {
	conn, err := db.ConnectionFn(DatabaseURL, db.WithDriver(DriverName))()
	if err != nil {
		return nil, err
	}

	sqlxDB := sqlx.NewDb(conn, DriverName)

	return sqlxDB, nil
}
