package db

import (
	"context"
	"testing"

	"github.com/MogLuiz/go-bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createdAccount := createRandomAccount(t)
	createRandomEntry(t, createdAccount)
}

func TestGetEntry(t *testing.T) {
	createdAccount := createRandomAccount(t)
	createdEntry := createRandomEntry(t, createdAccount)

	fetchedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedEntry)

	require.Equal(t, createdEntry.ID, fetchedEntry.ID)
	require.Equal(t, createdEntry.AccountID, fetchedEntry.AccountID)
	require.Equal(t, createdEntry.Amount, fetchedEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, fetchedEntry.CreatedAt, 0)
}

func TestListEntries(t *testing.T) {
	createdAccount := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, createdAccount)
	}

	args := ListEntriesParams{
		AccountID: createdAccount.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, args.AccountID, entry.AccountID)
	}
}
