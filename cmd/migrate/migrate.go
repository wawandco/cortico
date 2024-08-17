package main

import (
	"fmt"

	"cortico/internal"
	"cortico/internal/migrations"

	"github.com/leapkit/leapkit/core/db"
	// postgres driver
	_ "github.com/lib/pq"
)

// The migrate command is used to ship our application
// with the latest database schema migrator. which can be invoked
// by running `migrate`.
func main() {
	conn, err := db.ConnectionFn(internal.DatabaseURL, db.WithDriver(internal.DriverName))()
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
	}

	err = db.RunMigrations(migrations.All, conn)
	if err != nil {
		fmt.Println("Error running migrations: ", err)
	}

	fmt.Println("✅ Migrations ran successfully")
}
