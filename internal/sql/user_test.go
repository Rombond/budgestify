package db_sql_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Rombond/budgestify/internal/password"
	db_sql "github.com/Rombond/budgestify/internal/sql"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	// Démarrer un conteneur MySQL temporaire
	mysqlContainer, err := mysql.Run(ctx,
		"mysql:9.4.0",
		mysql.WithDatabase("testdb"),
		mysql.WithUsername("root"),
		mysql.WithPassword("password"),
	)
	if err != nil {
		t.Fatalf("Failed to start MySQL container: %v", err)
	}
	defer mysqlContainer.Terminate(ctx)

	// Obtenir l'URL de connexion
	connStr, err := mysqlContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatalf("Failed to get connection string: %v", err)
	}

	// Connecter à la base de données
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	defer db.Close()

	// Initialiser la base de données
	db_sql.InitDatabase(db)

	// Test
	hash, err := password.ParamToByte("ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff")
	if err != nil {
		t.Fatalf("Failed to convert password hash: %v", err)
	}
	id, err := db_sql.CreateUser(db, "Test User", "test", hash)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	if id <= 0 {
		t.Fatalf("Expected user ID to be greater than 0, got %d", id)
	}
	t.Logf("User created with ID: %d", id)
}
