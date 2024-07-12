package conf

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDb() {
	if db == nil {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			panic(err)
		}

		os.MkdirAll(homeDir+"/.data/daily_log", 0755)
		if _, err := os.Open(homeDir + "/.data/daily_log/data.db"); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("Creating database file in " + homeDir + "/.data/daily_log/data.db")
				os.Create(homeDir + "/.data/daily_log/data.db")
			} else {
				panic(err)
			}
		}

		db, err = sql.Open("sqlite3", homeDir+"/.data/daily_log/data.db")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// defer db.Close()

		if !tableExists("log_entries") {
			_, err = db.Exec("CREATE TABLE `log_entries` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `project` VARCHAR(255) NOT NULL, `task` VARCHAR(255) NULL, `hours` FLOAT(5) NOT NULL DEFAULT 0.0, `date` DATE NOT NULL)")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

func tableExists(tableName string) bool {

	_, tableExists := db.Query("SELECT * FROM " + tableName)
	return tableExists == nil
}

func Get() *sql.DB {
	initDb()
	return db
}

func Close() {
	if db != nil {
		db.Close()
	}
}
