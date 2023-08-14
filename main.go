package main

import (
	"database/sql"
	"fmt"
	// "log"
	// _ "github.com/lib/pq"
)

var db *sql.DB

type Data struct {
	ingredient_name        string
	ingredient_quantity    int
	quantity_type          string
	nutrient_quantity      int
	nutrient_quantity_type string
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	// password = ""
	dbname = "golang_practice"
	table  = "food_details"
	// column1 = "videogame_title"
	// column2 = "videogame_platform"
)

func main() {
	// Connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	// Open database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	// CheckError("Error opening database:", err)
	if err != nil {
		fmt.Println("Error opening database:")
		return
	}

	// Close database
	defer db.Close()

	// Check DB
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:")
		return
	}

	fmt.Println("Connected!")

	// Additional debug statements
	//fmt.Println("Database connection details:", psqlconn)

	// Dynamic Insertion of Data
	insertData := `INSERT INTO "food_details" ("ingredient_name", "ingredient_quantity", "quantity_type", "nutrient_quantity", "nutrient_quantity_type") VALUES($1, $2, $3, $4, $5)`
	// var videogameID uint32
	_, err = db.Exec(insertData, "Spider-Man 2", "Playstation 2", "June 28, 2004", "Activision")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data insertion successful!")

	// Identify duplicate and delete records
	// duplicateQuery := fmt.Sprintf(`
	// 	WITH duplicates AS (
	// 		SELECT %s, %s, COUNT(*) AS cnt
	// 		FROM %s
	// 		GROUP BY %s, %s
	// 		HAVING COUNT(*) > 1
	// 	)
	// 	DELETE FROM %s
	// 	WHERE (%s, %s) IN (SELECT %s, %s FROM duplicates);
	// 	`, column1, column2, table, column1, column2, table, column1, column2, column1, column2)

	// // Execute the delete query
	// _, err = db.Exec(duplicateQuery)
	// if err != nil {
	// 	log.Fatal("Error deleting duplicate records:", err)
	// }

	// fmt.Println("Duplicate records deleted successfully!")

	// Query Table
	rows, err := db.Query(fmt.Sprintf(`SELECT "ingredient_name", "ingredient_quantity" FROM %s`, table))
	if err != nil {
		fmt.Println("Error executing query")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var videogame_title string
		var videogame_platform string

		err = rows.Scan(&videogame_title, &videogame_platform)
		if err != nil {
			fmt.Println("Error executing query")
			return
		}

		fmt.Println(videogame_title, videogame_platform)
	}
}
