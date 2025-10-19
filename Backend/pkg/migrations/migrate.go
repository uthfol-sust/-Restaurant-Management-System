package migrations

import (
	"database/sql"
	"errors"
	"log"
)


func Migrate(db *sql.DB) error {
	if db == nil {
		return errors.New("migrations: nil db provided")
	}
	
	UsersTable:= `CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		password VARCHAR(250),
		email VARCHAR(50),
		phone_no VARCHAR(15),
		role VARCHAR(20) DEFAULT 'Admin',
		shift_time VARCHAR(50),
		join_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(UsersTable)
	if err != nil{
		log.Fatalf("Failed to create users table: %v", err)
	}


	log.Println("migrations: no models registered for auto-migrate")
	return nil
}