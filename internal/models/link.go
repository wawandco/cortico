package models

import (
	"cmp"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/itchyny/base58-go"
)

type Link struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Original  string    `db:"original" json:"original"`
	Short     string    `db:"short" json:"short"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// LinksService is the interface that wraps the basic CRUD operations
// for the Link model
type LinksService interface {
	Create(link *Link) error
	Find(shortUrl string) (Link, error)
}

func (link *Link) GenerateShortLink() error {
	urlHashBytes := sha256Of(link.Original)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return err
	}

	link.Short = finalString[:8]
	return nil
}

func sha256Of(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))

	return hash.Sum(nil)
}

func base58Encoded(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (link *Link) FullLink() string {
	base := cmp.Or(os.Getenv("BASE_URL"), "http://localhost:3000")
	return fmt.Sprintf("%v/%v", base, link.Short)
}
