package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
	"database/sql"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	assert.NoError(t, err)
	clientID := 1
	client, err := selectClient(db, clientID)
	assert.NoError(t, err)
	assert.Equal(t, clientID, client.ID)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	assert.NoError(t, err)
	clientID := -1
	_, err = selectClient(db, clientID)
	require.ErrorIs(t, err, sql.ErrNoRows,  "ожидалась ошибка sql.ErrNoRows")
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	assert.NoError(t, err)
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	cl.ID, err = insertClient(db, cl)
	assert.NoError(t, err)
	assert.NotEmpty(t, cl.ID)
	_, err = selectClient(db, cl.ID)
	assert.NoError(t, err)
	
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	assert.NoError(t, err)
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)
	_, err = selectClient(db, cl.ID)
	require.NoError(t, err)
	err = deleteClient(db, cl.ID)
	require.NoError(t, err)
	_, err = selectClient(db, cl.ID)
	require.ErrorIs(t, err, sql.ErrNoRows, "Ожидалась ошибка sql.ErrNoRows")
}
