package links

import (
	"cortico/internal/models"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/leapkit/core/form"
	"github.com/leapkit/leapkit/core/render"
)

func ShortURL(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	linkService := r.Context().Value("links").(models.LinksService)

	link := models.Link{}
	if err := form.Decode(r, &link); err != nil {
		rw.Set("error", "Internal issue generating the URL, try later.")
		slog.Error("Error parsing the form", "err", err.Error())
	}

	if err := link.ValidateURL(); err != nil {
		rw.Set("error", "Please enter a valid URL")
		slog.Error("Error validating the URL", "err", err.Error())
	}

	if err := link.GenerateShortLink(); err != nil {
		rw.Set("error", "Internal issue generating the URL, try later.")
		slog.Error("Error generating the short URL", "err", err.Error())
	}

	if err := linkService.Create(&link); err != nil {
		rw.Set("error", "Internal issue generating the URL, try later.")
		slog.Error("Error storing the link", "err", err.Error())
	}

	rw.Set("shortUrl", link.FullLink())
	err := rw.RenderClean("links/result.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	linkService := r.Context().Value("links").(models.LinksService)
	shortUrl := r.PathValue("short_url")

	link, err := linkService.Find(shortUrl)
	if err == sql.ErrNoRows {
		err := rw.RenderClean("links/404.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, link.Original, http.StatusSeeOther)
}
