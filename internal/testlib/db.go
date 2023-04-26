package testlib

import (
	"github.com/stretchr/testify/require"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"testing"
)

// NewTestDB creates a new in-memory SQLite database instance for testing purposes.
func NewTestDB(t *testing.T) *db.DataBase {
	// Create a new database connection
	testDb, err := db.NewDataBase(nil, db.Config{FileName: ":memory:"})
	require.NoErrorf(t, err, "Test database creation failed: %q", err)

	// Initialize the schema
	err = testDb.InitialMigration()
	require.NoErrorf(t, err, "Test database InitialMigration failed: %q", err)

	return testDb
}
