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


	ProductTable := `CREATE TABLE IF NOT EXISTS products(
	    product_id INT PRIMARY KEY,
		product_name VARCHAR(100) NOT NULL,
		category VARCHAR(50) NOT NULL,
		price DECIMAL(10,2) NOT NULL,
		availability_status VARCHAR(30) DEFAULT 'Out of Stock'
	)`

	_, err = db.Exec(ProductTable)
	if err != nil{
		log.Fatalf("Failed to create product table: %v", err)
	}



    InventoryTable := `CREATE TABLE IF NOT EXISTS inventory (
		id INT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		stock DECIMAL(10,2) NOT NULL,
		unit VARCHAR(20) NOT NULL,
		level DECIMAL(10,2) NOT NULL,
		last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(InventoryTable)
	if err != nil {
		log.Fatalf("❌ Failed to create inventory table: %v", err)
	}



	invent_pro := `CREATE TABLE IF NOT EXISTS inventory_product (
						product_id INT,
						id INT,
						quantity_required INT NOT NULL,
						FOREIGN KEY (product_id) REFERENCES products(product_id)
							ON UPDATE CASCADE
							ON DELETE CASCADE,
						FOREIGN KEY (id) REFERENCES inventory(id)
							ON UPDATE CASCADE
							ON DELETE CASCADE
					);`
	 
	_, err = db.Exec(invent_pro)
	if err != nil {
		log.Fatalf("❌ Failed to create inventory_product table: %v", err)
	}

	supplier := `CREATE TABLE IF NOT EXISTS suppliers (
					supplier_id INT PRIMARY KEY,
					name VARCHAR(100) NOT NULL,
					contact_no VARCHAR(20),
					email VARCHAR(100),
					address TEXT
				);`
	
	_, err = db.Exec(supplier)
	if err != nil {
		log.Fatalf("❌ Failed to create supplier table: %v", err)
	}


	log.Println("✅ Database migration completed successfully!")
	return nil
}


