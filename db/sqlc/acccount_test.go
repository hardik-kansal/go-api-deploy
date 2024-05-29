package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/hardik-kansal/go-api-deploy/utils"
)
func Test_account(t *testing.T) {
	arg:= CreateAccountParams{
		Owner:util.RandomOwner(),
		Currency: util.RandomCurrency(),
		Balance:  util.RandomMoney(),
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
