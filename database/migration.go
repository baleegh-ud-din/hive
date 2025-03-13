package database

func CreateMigrations() error {
	migrations := []struct {
		version        int
		name           string
		migrationquery string
	}{
		// {
		// 	1,
		// 	"offices",
		// 	`
		// 	DROP COLUMN image_type,
		// 	DROP COLUMN image;
		// 	DROP COLUMN image_url;
		// 	`,
		// },
		// Add more migrations here...
	}

	for _, migration := range migrations {
		if err := CreateMigration(migration.version, migration.name, migration.migrationquery); err != nil {
			return err
		}
	}
	return nil
}
