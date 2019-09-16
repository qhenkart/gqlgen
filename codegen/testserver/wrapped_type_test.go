package testserver

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/qhenkart/gqlgen/client"
	"github.com/qhenkart/gqlgen/codegen/testserver/otherpkg"
	"github.com/qhenkart/gqlgen/handler"
	"github.com/stretchr/testify/require"
)

func TestWrappedTypes(t *testing.T) {
	resolvers := &Stub{}

	srv := httptest.NewServer(handler.GraphQL(NewExecutableSchema(Config{Resolvers: resolvers})))
	c := client.New(srv.URL)

	resolvers.QueryResolver.WrappedScalar = func(ctx context.Context) (scalar WrappedScalar, e error) {
		return WrappedScalar("hello"), nil
	}

	resolvers.QueryResolver.WrappedStruct = func(ctx context.Context) (wrappedStruct *WrappedStruct, e error) {
		wrapped := WrappedStruct(otherpkg.Struct{
			Name: "hello",
		})
		return &wrapped, nil
	}

	t.Run("wrapped struct", func(t *testing.T) {
		var resp struct {
			WrappedStruct struct {
				Name string
			}
		}

		err := c.Post(`query { wrappedStruct { name } }`, &resp)
		require.NoError(t, err)

		require.Equal(t, "hello", resp.WrappedStruct.Name)
	})

	t.Run("wrapped scalar", func(t *testing.T) {
		var resp struct {
			WrappedScalar string
		}

		err := c.Post(`query { wrappedScalar }`, &resp)
		require.NoError(t, err)

		require.Equal(t, "hello", resp.WrappedScalar)
	})
}
