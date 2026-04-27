package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestNextStatus(t *testing.T) {
	// prepare
	db, err := sql.Open("sqlite", "tracker.db")
	require.NoError(t, err)
	defer db.Close()

	store := NewParcelStore(db)
	service := NewParcelService(store)
	parcel := getTestParcel()
	// обновите статус, убедитесь в отсутствии ошибки

	id, err := store.Add(parcel)
	require.NoError(t, err)
	require.NotZero(t, id)
	// добавьте новую посылку в БД, убедитесь в отсутствии ошибки и наличии идентификатора

	err = service.NextStatus(id)
	require.NoError(t, err)
	// check
	p, err := store.Get(id)
	require.NoError(t, err)
	assert.Equal(t, p.Status, ParcelStatusSent)
	// получите добавленную посылку и убедитесь, что статус обновился
}
