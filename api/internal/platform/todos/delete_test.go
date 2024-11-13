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

func TestDelete(t *testing.T) {

	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: "123",
		},
	}
	ctx := context.WithValue(context.Background(), "key", "value")
	contextMatch := mock.MatchedBy(func(ctx context.Context) bool {
		return ctx.Value("key") == "value"
	})

	deleteItemMatch := mock.MatchedBy(func(input *dynamodb.DeleteItemInput) bool {
		return input.TableName != nil &&
			*input.TableName == "testTable" &&
			reflect.DeepEqual(input.Key, key)
	})

	optsMatch := mock.MatchedBy(func(opts []func(*dynamodb.Options)) bool {
		return len(opts) == 0
	})

	t.Run("successful delete", func(t *testing.T) {
		mockDynamoDB := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDynamoDB,
			tableName: "testTable",
		}
		mockDynamoDB.On("DeleteItem", contextMatch, deleteItemMatch, optsMatch).
			Return(&dynamodb.DeleteItemOutput{}, nil)

		err := repo.Delete(ctx, "123")
		assert.NoError(t, err)

	})

	t.Run("delete item error", func(t *testing.T) {
		mockDynamoDB := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDynamoDB,
			tableName: "testTable",
		}

		mockDynamoDB.On("DeleteItem", contextMatch, deleteItemMatch, optsMatch).
			Return(nil, errors.ErrUnsupported)

		err := repo.Delete(ctx, "123")
		assert.Error(t, err)
		assert.ErrorIs(t, err, errors.ErrUnsupported)
	})
}
