package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "io"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Data struct {
	IngredientName       string  //`json:"ingredient_name"`
	IngredientQuantity   float64 //`json:"ingredient_quantity"`
	QuantityType         string  //`json:"quantity_type"`
	NutrientQuantity     int     //`json:"nutrient_quantity"`
	NutrientQuantityType string  //`json:"nutrient_quantity_type"`
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	// password = ""
	dbname  = "golang_practice"
	table   = "food_details"
	column1 = "ingredient_name"
	column2 = "ingredient_quantity"
)

func main() {
	// Connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	// Connect to database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	// CheckError("Error opening database:", err)
	if err != nil {
		fmt.Println("Error opening database:", err)
		fmt.Println("Database connection details:", psqlconn)
		return
	}

	// Close database
	defer db.Close()

	// Check DB
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:")
		return
		//fmt.Println("Database connection details:", psqlconn)
	}

	fmt.Println("Connected!")

	// Ingestion of JSON Extracts (Pause for now)
	file_location := "./milk_and_dairy_sample.json"
	// file_location := "./sample.json"

	jsonFile, err := os.Open(file_location)
	if err != nil {
		log.Fatal("Error when opening file:", err)
		return
	}
	defer jsonFile.Close()

	fmt.Println("Successfully opened file")

	// Read the file content into a []byte slice using os.ReadFile
	fileContent, err := os.ReadFile(file_location)
	if err != nil {
		log.Fatal("Error reading file:", err)
		return
	}
	fmt.Println(string(fileContent))

	// Unmarshall data into 'payload'
	var payload []Data
	err = json.Unmarshal(fileContent, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshall():", err)
	}

	// Print unmarshalled data
	for _, data := range payload {
		log.Printf("Ingredient: %s\n", data.IngredientName)
		log.Printf("Ingredient Quantity: %f\n", data.IngredientQuantity)
		log.Printf("Quantity Type: %s\n", data.QuantityType)
		log.Printf("Nutrient Quantity: %d\n", data.NutrientQuantity)
		log.Printf("Nutrient Quantity Type: %s\n", data.NutrientQuantityType)
	}
	// Read opened jsonFile as a byte array
	// byteValue, _ := io.ReadAll(jsonFile)

	// Initialize Data Array (Second attempt)

	// jsonData := `
	// 		[{
	// 			"ingredient_name": "Sesame seeds",
	// 			"ingredient_quantity": 0.25,
	// 			"quantity_type": "cup",
	// 			"nutrient_quantity": 351,
	// 			"nutrient_quantity_type": "milligrams"
	// 		},

	// 		{
	// 			"ingredient_name": "Sardines (with bones)",
	// 			"ingredient_quantity": 3.75,
	// 			"quantity_type": "ounce-can",
	// 			"nutrient_quantity": 351,
	// 			"nutrient_quantity_type": "milligrams"
	// 		}]`

	// var data []map[string]interface{}
	// ingestion_err := json.Unmarshal([]byte(jsonData), &data)
	// if ingestion_err != nil {
	// 	fmt.Printf("Could not unmarshall json: %s\n", ingestion_err)
	// 	return
	// }

	// fmt.Printf("Json mapped: %v\n", data)

	// Dynamic Insertion of Data
	manual_data_insertion := false
	if manual_data_insertion {
		insertData := `INSERT INTO "food_details" ("ingredient_name", "ingredient_quantity", "quantity_type", "nutrient_quantity", "nutrient_quantity_type") VALUES($1, $2, $3, $4, $5)`
		// var videogameID uint32
		_, err = db.Exec(insertData, "Spider-Man 2", "Playstation 2", "June 28, 2004", "Activision")
		if err != nil {
			panic(err)
		}
		fmt.Println("Data insertion successful!")
	}

	duplicate_search := false
	if duplicate_search {
		// Identify duplicate and delete records
		duplicateQuery := fmt.Sprintf(`
			WITH duplicates AS (
				SELECT %s, %s, COUNT(*) AS cnt
				FROM %s
				GROUP BY %s, %s
				HAVING COUNT(*) > 1
			)
			DELETE FROM %s
			WHERE (%s, %s) IN (SELECT %s, %s FROM duplicates);
			`, column1, column2, table, column1, column2, table, column1, column2, column1, column2)

		// Execute the delete query
		_, err = db.Exec(duplicateQuery)
		if err != nil {
			log.Fatal("Error deleting duplicate records:", err)
		}

		fmt.Println("Duplicate records deleted successfully!")
	}

	// Query Table
	query_table := false
	if query_table {
		rows, err := db.Query(fmt.Sprintf(`SELECT "ingredient_name", "ingredient_quantity" FROM %s`, table))
		if err != nil {
			fmt.Println("Error executing query")
			return
		}
		defer rows.Close()

		for rows.Next() {
			var ingredient_name string
			var ingredient_quantity string

			err = rows.Scan(&ingredient_name, &ingredient_quantity)
			if err != nil {
				fmt.Println("Error executing query")
				return
			}

			fmt.Println(ingredient_name, ingredient_quantity)
		}
	}

}
