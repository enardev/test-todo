package todos

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAll(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	contextMatch := mock.MatchedBy(func(ctx context.Context) bool {
		return ctx.Value("key") == "value"
	})

	inputMatch := mock.MatchedBy(func(input *dynamodb.ScanInput) bool {
		return input.TableName != nil &&
			*input.TableName == "testTable"
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
		mockDBClient.On("Scan", contextMatch, inputMatch, optsMatch).
			Return(&dynamodb.ScanOutput{
				Items: []map[string]types.AttributeValue{
					{
						"id":    &types.AttributeValueMemberS{Value: "1"},
						"title": &types.AttributeValueMemberS{Value: "Test ToDo 1"},
					},
					{
						"id":    &types.AttributeValueMemberS{Value: "2"},
						"title": &types.AttributeValueMemberS{Value: "Test ToDo 2"},
					},
				},
			}, nil)

		toDos, err := repo.FindAll(ctx)
		assert.NoError(t, err)
		assert.Len(t, toDos, 2)
		assert.Equal(t, "Test ToDo 1", toDos[0].Title)
		assert.Equal(t, "Test ToDo 2", toDos[1].Title)
	})

	t.Run("scan error", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}
		mockDBClient.On("Scan", contextMatch, inputMatch, optsMatch).
			Return(nil, errors.ErrUnsupported)

		toDos, err := repo.FindAll(ctx)
		assert.ErrorIs(t, err, errors.ErrUnsupported)
		assert.Empty(t, toDos)
	})

	t.Run("unmarshal error", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}
		mockDBClient.On("Scan", contextMatch, inputMatch, optsMatch).
			Return(&dynamodb.ScanOutput{
				Items: []map[string]types.AttributeValue{
					{
						"id":        &types.AttributeValueMemberS{Value: "1"},
						"completed": &types.AttributeValueMemberN{Value: "1.24"},
					},
				},
			}, nil)

		toDos, err := repo.FindAll(ctx)
		assert.Error(t, err)
		assert.Empty(t, toDos)
	})
}
