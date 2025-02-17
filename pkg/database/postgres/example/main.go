package example

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	database "crypto-dashboard/pkg/database/postgres"
// )

// func Init() {
// 	dbConfig := database.ConnectionConfig()
// 	conn, err := database.NewConnection(dbConfig)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}
// 	defer func() {
// 		if err := conn.Close(); err != nil {
// 			log.Printf("Failed to close the connection: %v", err)
// 		}
// 	}()

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := conn.HealthCheck(ctx); err != nil {
// 		log.Fatalf("Health check failed: %v", err)
// 	}

// 	fmt.Println("Database connection successful!")

// 	if err := createUserTable(ctx, conn); err != nil {
// 		log.Fatalf("Failed to create user table: %v", err)
// 	}
// 	fmt.Println("User table created!")
// }

// func createUserTable(ctx context.Context, conn *database.Connection) error {
// 	query := `
// 		CREATE TABLE IF NOT EXISTS users (
// 			id SERIAL PRIMARY KEY,
// 			name VARCHAR(100) NOT NULL,
// 			email VARCHAR(100) UNIQUE NOT NULL,
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		);
// 	`
// 	_, err := conn.DB().Exec(ctx, query)
// 	if err != nil {
// 		return fmt.Errorf("failed to execute query: %w", err)
// 	}
// 	return nil
// }
