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

func TestExists(t *testing.T) {
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

	t.Run("Item exists", func(t *testing.T) {
		mockClient := new(mockDynamoDBClient)
		s := &repo{
			dbClient:  mockClient,
			tableName: "testTable",
		}

		successOutput := &dynamodb.GetItemOutput{
			Item: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{
					Value: "123",
				},
			},
		}

		mockClient.On("GetItem", contextMatch, inputMatch, optsMatch).
			Return(successOutput, nil)

		exists, err := s.Exists(ctx, "123")
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("Item does not exist", func(t *testing.T) {
		mockClient := new(mockDynamoDBClient)
		s := &repo{
			dbClient:  mockClient,
			tableName: "testTable",
		}

		successOutput := &dynamodb.GetItemOutput{}

		mockClient.On("GetItem", contextMatch, inputMatch, optsMatch).
			Return(successOutput, nil)

		exists, err := s.Exists(ctx, "123")
		assert.NoError(t, err)
		assert.False(t, exists)
	})

	t.Run("GetItem error", func(t *testing.T) {
		mockClient := new(mockDynamoDBClient)
		s := &repo{
			dbClient:  mockClient,
			tableName: "testTable",
		}

		mockClient.On("GetItem", contextMatch, inputMatch, optsMatch).
			Return(nil, errors.ErrUnsupported)

		exists, err := s.Exists(ctx, "123")
		assert.ErrorIs(t, err, errors.ErrUnsupported)
		assert.False(t, exists)
	})
}
