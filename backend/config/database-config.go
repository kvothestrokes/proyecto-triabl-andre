package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() *sql.DB {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// bin, err := ioutil.ReadFile("/run/secrets/db-password")
	// if err != nil {
	// 	fmt.Printf("Error al leer la password: %v", err.Error())
	// 	return nil
	// }
	db, err := sql.Open("mysql", fmt.Sprintf("root:tribalpass@tcp(db:3306)/biblioteca_canciones"))

	// fmt.Printf("\nConectando a la base de datos...\n %v \n", connectionString)
	// db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Printf("\nError al conectar a la base de datos: %v\n", err)
		panic(err.Error())
	}

	return db
}
