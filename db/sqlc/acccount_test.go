package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
)
func Test_account(t *testing.T) {
	arg:= CreateAccountParams{
		Owner:    "abc",
		Currency: "USD",
		Balance:  10,
	}
	insertedAccount, err := queries.CreateAccount(context.Background(),arg)
	require.NoError(t, err)
	require.NotEmpty(t, insertedAccount)
	require.Equal(t,insertedAccount.Owner,arg.Owner)
	require.Equal(t,insertedAccount.Balance,arg.Balance)
	require.Equal(t,insertedAccount.Currency,arg.Currency)
	require.NotZero(t,insertedAccount.ID)
	require.NotZero(t,insertedAccount.CreatedAt)
}
