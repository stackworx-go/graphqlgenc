package integration

import (
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stackworx-go/gqlgen-relay/internal/integration/graph"
	"github.com/stackworx-go/gqlgen-relay/internal/integration/graph/generated"
	"github.com/stackworx-go/gqlgen-relay/internal/integration/graph/model"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	// given
	resolvers := &graph.Resolver{}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	resolvers.Todos.Data = []*model.Todo{{
		ID:   "1",
		Text: "Buy Groceries",
		Done: false,
		User: &model.User{
			ID:   "1",
			Name: "John",
		},
	}}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	client := Client{
		Url: ts.URL,
	}

	// when
	data, err := client.TodosQuery()

	// then
	assert.NoError(t, err)
	assert.Equal(t, data, &TodosQueryPayload{
		Todos: nil,
	})
}
