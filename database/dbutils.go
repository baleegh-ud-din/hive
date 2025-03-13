package database

import (
	"fmt"
	"strings"
)

func CreateEnum(enumName string, values []string) error {
	enumValues := "'" + strings.Join(values, "', '") + "'"
	query := fmt.Sprintf(`
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = '%s') THEN
                CREATE TYPE %s AS ENUM (%s);
            END IF;
        END $$;
    `, enumName, enumName, enumValues)

	_, err := DB.Exec(query)
	if err != nil {
		return fmt.Errorf("❌ Error creating enum %s: %w", enumName, err)
	}
	msg := fmt.Sprintf("✅ %s ENUM is ensured.", enumName)
	logger.Success(msg)
	return nil
}

func CreateTable(tableName, schema string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, schema)
	_, err := DB.Exec(query)
	if err != nil {
		return fmt.Errorf("❌ Error creating table %s: %w", tableName, err)
	}
	msg := fmt.Sprintf("✅ %s TABLE is ensured.", tableName)
	logger.Success(msg)
	return nil
}

func CreateIndex(indexName, tableName, column string) error {
	query := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (%s);",
		indexName, tableName, column)

	_, err := DB.Exec(query)
	if err != nil {
		return fmt.Errorf("❌ Error creating index %s on %s(%s): %w", indexName, tableName, column, err)
	}
	msg := fmt.Sprintf("✅ %s INDEX is ensured on %s(%s).", indexName, tableName, column)
	logger.Success(msg)
	return nil
}

func CreateMigration(version int, tableName, migration string) error {
	// Start a transaction
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("❌ Error starting transaction: %w", err)
	}
	defer tx.Rollback()

	// Create migrations table if it doesn't exist
	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			version INT NOT NULL UNIQUE,
			applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return fmt.Errorf("❌ creating migrations table: %w", err)
	}

	// Check if migration has already been applied
	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM migrations WHERE version = $1)", version).Scan(&exists)
	if err != nil {
		return fmt.Errorf("❌ Error checking if migration exists: %w", err)
	}

	if exists {
		logger.Info(fmt.Sprintf("🛩️ Migration version %d has already been applied.", version))
		return nil
	}

	// Apply the migration
	query := fmt.Sprintf("ALTER TABLE %s %s", tableName, migration)
	_, err = tx.Exec(query)
	if err != nil {
		msg := fmt.Sprintf("❌ Error migrating table %s: %v", tableName, err)
		logger.Error(msg)
	}

	// Record the migration
	_, err = tx.Exec("INSERT INTO migrations (version) VALUES ($1)", version)
	if err != nil {
		msg := fmt.Sprintf("❌ Error recording migration version %d: %v", version, err)
		logger.Error(msg)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		msg := fmt.Sprintf("❌ Error committing migration: %v", err)
		logger.Error(msg)
	}

	logger.Success(fmt.Sprintf("🛫 %s TABLE is migrated to version %d.", tableName, version))
	return nil
}
