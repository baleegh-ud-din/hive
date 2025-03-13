package database

func CreateSchemas() error {
	// ENUM
	enums := []struct {
		name   string
		values []string
	}{
		// {"approval_status", []string{"Pending", "Approved", "Requesting Change", "Rejected"}},
	}

	for _, enum := range enums {
		if err := CreateEnum(enum.name, enum.values); err != nil {
			return err
		}
	}

	// TABLES
	tables := []struct {
		name   string
		schema string
	}{
		// {"offices", `
		// 	code varchar(20) PRIMARY KEY,
		// 	registered_name VARCHAR(255) NOT NULL,
		// 	office_type office_type NOT NULL,
		// 	office_status office_status NOT NULL,
		// 	sys_poc VARCHAR(20) NOT NULL,
		// 	net_poc VARCHAR(20) NOT NULL,
		// 	dc_poc VARCHAR(20) NOT NULL,
		// 	admin_poc VARCHAR(20) NOT NULL,
		// 	address VARCHAR(500) NOT NULL,
		// 	city VARCHAR(50) NOT NULL,
		// 	country VARCHAR(50) NOT NULL,
		// 	zip VARCHAR(20) NOT NULL,
		// 	phone VARCHAR(20) NOT NULL,
		// 	email VARCHAR(50) NOT NULL,
		// 	description TEXT NOT NULL,
		// 	last_urid BIGINT NOT NULL,
		// 	approved_by VARCHAR(20) NOT NULL,
		// 	last_approved_by VARCHAR(20) NOT NULL,
		// 	created_by VARCHAR(20) NOT NULL,
		// 	updated_by VARCHAR(20) NOT NULL,
		// 	created_at BIGINT DEFAULT EXTRACT(EPOCH FROM now())::BIGINT,
		// 	updated_at BIGINT DEFAULT EXTRACT(EPOCH FROM now())::BIGINT
		// `},
	}

	for _, table := range tables {
		if err := CreateTable(table.name, table.schema); err != nil {
			return err
		}
	}

	// INDEXES
	indexes := []struct {
		name      string
		tableName string
		column    string
	}{
		// {"idx_approvals_urid", "approvals", "urid"},
	}

	for _, index := range indexes {
		if err := CreateIndex(index.name, index.tableName, index.column); err != nil {
			return err
		}
	}

	return nil
}
