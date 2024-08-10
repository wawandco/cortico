package links

import (
	"cortico/internal/models"
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

	// TODO: Validate URL
	// Should it generate a new token for the same URL?
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
	linkService := r.Context().Value("links").(models.LinksService)
	shortUrl := r.PathValue("short_url")

	link, err := linkService.Find(shortUrl)
	if err != nil {
		// TODO: Link Not Found
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, link.Original, http.StatusSeeOther)
}
