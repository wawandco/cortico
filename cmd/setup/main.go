package main

import (
	"fmt"
	"os"

	"os/exec"

	"cortico/internal"
	"cortico/internal/migrations"

	"github.com/leapkit/leapkit/core/db"

	// Load environment variables
	_ "github.com/leapkit/leapkit/core/tools/envload"
	// postgres driver
	_ "github.com/lib/pq"
)

func main() {
	// Running the tailo setup command
	cmd := exec.Command("go", "run", "github.com/paganotoni/tailo/cmd/build@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Tailwind CSS setup successfully")
	err = db.Create(internal.DatabaseURL)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Database created successfully")
	conn, err := db.ConnectionFn(internal.DatabaseURL, db.WithDriver(internal.DriverName))()
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
	}

	err = db.RunMigrations(migrations.All, conn)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Migrations ran successfully")
}
