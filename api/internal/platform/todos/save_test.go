package todos

import (
	"context"
	"errors"
	"reflect"
	"test-todo/api/internal/domain/todos"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSave(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	contextMatch := mock.MatchedBy(func(ctx context.Context) bool {
		return ctx.Value("key") == "value"
	})

	optsMatch := mock.MatchedBy(func(opts []func(*dynamodb.Options)) bool {
		return len(opts) == 0
	})

	date := time.Date(2024, time.April, 21, 22, 05, 16, 0, time.UTC)
	toDo := todos.ToDo{
		ID:        "1",
		Title:     "Test ToDo",
		UpdatedAt: date,
	}

	item := map[string]types.AttributeValue{
		"id":         &types.AttributeValueMemberS{Value: "1"},
		"title":      &types.AttributeValueMemberS{Value: "Test ToDo"},
		"completed":  &types.AttributeValueMemberBOOL{Value: false},
		"created_at": &types.AttributeValueMemberS{Value: "0001-01-01T00:00:00Z"},
		"updated_at": &types.AttributeValueMemberS{Value: date.Format(time.RFC3339)},
	}

	putItemMatch := mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
		return input.TableName != nil &&
			*input.TableName == "testTable" &&
			input.Item != nil &&
			reflect.DeepEqual(input.Item, item)
	})

	t.Run("successful save", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}
		mockDBClient.On("PutItem", contextMatch, putItemMatch, optsMatch).
			Return(&dynamodb.PutItemOutput{}, nil)

		err := repo.Save(ctx, toDo)
		assert.NoError(t, err)
	})

	t.Run("dynamodb put item error", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}
		mockDBClient.On("PutItem", contextMatch, putItemMatch, optsMatch).
			Return(nil, errors.New("dynamodb error"))

		err := repo.Save(ctx, toDo)
		assert.Error(t, err)
	})
}
