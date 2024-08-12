package links

import (
	"cortico/internal/models"
	"database/sql"
	"net/http"

	"github.com/leapkit/core/form"
	"github.com/leapkit/leapkit/core/render"
)

func ShortURL(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	linkService := r.Context().Value("links").(models.LinksService)

	link := models.Link{}
	if err := form.Decode(r, &link); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := link.ValidateURL(); err != nil {
		rw.Set("error", "Please enter a valid URL")

		err := rw.Render("home/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	// TODO
	// Expiration time

	if err := link.GenerateShortLink(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := linkService.Create(&link); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rw.Set("shortUrl", link.FullLink())
	err := rw.Render("home/index.html")
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
		err := rw.RenderClean("home/404.html")
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
