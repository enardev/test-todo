package todos

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindByID(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	contextMatch := mock.MatchedBy(func(ctx context.Context) bool {
		return ctx.Value("key") == "value"
	})
	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: "123",
		},
	}
	inputMatch := mock.MatchedBy(func(input *dynamodb.GetItemInput) bool {
		return input.TableName != nil &&
			*input.TableName == "testTable" &&
			reflect.DeepEqual(input.Key, key)
	})
	optsMatch := mock.MatchedBy(func(opts []func(*dynamodb.Options)) bool {
		return len(opts) == 0
	})

	t.Run("success", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}

		mockDBClient.On("GetItem", contextMatch, inputMatch, optsMatch).
			Return(&dynamodb.GetItemOutput{
				Item: map[string]types.AttributeValue{
					"id":    &types.AttributeValueMemberS{Value: "123"},
					"title": &types.AttributeValueMemberS{Value: "Test ToDo"},
				},
			}, nil)

		toDo, err := repo.FindByID(ctx, "123")
		assert.NoError(t, err)
		assert.Equal(t, "123", toDo.ID)
		assert.Equal(t, "Test ToDo", toDo.Title)
	})

	t.Run("item not found", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}
		mockDBClient.On("GetItem", contextMatch, inputMatch, optsMatch).
			Return(&dynamodb.GetItemOutput{}, nil)

		_, err := repo.FindByID(ctx, "123")
		assert.Error(t, err)
		assert.Equal(t, "item not found", err.Error())
	})

	t.Run("dynamodb error", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}
		mockDBClient.On("GetItem", contextMatch, inputMatch, optsMatch).
			Return(nil, errors.New("dynamodb error"))

		_, err := repo.FindByID(ctx, "123")
		assert.Error(t, err)
		assert.Equal(t, "dynamodb error", err.Error())
	})

	t.Run("unmarshal error", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}
		mockDBClient.On("GetItem", contextMatch, inputMatch, optsMatch).
			Return(&dynamodb.GetItemOutput{
				Item: map[string]types.AttributeValue{
					"id":        &types.AttributeValueMemberS{Value: "123"},
					"completed": &types.AttributeValueMemberN{Value: "1.24"},
				},
			}, nil)

		_, err := repo.FindByID(ctx, "123")
		assert.Error(t, err)
	})
}
