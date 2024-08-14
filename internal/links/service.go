package links

import (
	"cortico/internal/models"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) models.LinksService {
	return &service{db: db}
}

func (service *service) Create(link *models.Link) error {
	link.ID = uuid.Must(uuid.NewV4())
	link.CreatedAt = time.Now()
	_, err := service.db.NamedExec(`INSERT INTO links (id, original, short, created_at) VALUES (:id, :original, :short, :created_at)`, link)
	return err
}

func (service *service) Find(shortUrl string) (models.Link, error) {
	link := models.Link{}
	err := service.db.Get(&link, `SELECT * FROM links WHERE short = $1 ORDER BY created_at DESC LIMIT 1`, shortUrl)
	return link, err
}
