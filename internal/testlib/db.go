package testlib

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/volmexfinance/relayers/relayer-srv/db"
)

// NewTestDB creates a new in-memory SQLite database instance for testing purposes.
func NewTestDB(t *testing.T) *db.DataBase {
	// Create a new database connection
	testDb, err := db.InitialMigration("efe", *logrus.New())
	require.NoErrorf(t, err, "Test database creation failed: %q", err)

	// Initialize the schema
	// err = testDb.InitialMigration()
	require.NoErrorf(t, err, "Test database InitialMigration failed: %q", err)

	return testDb
}
