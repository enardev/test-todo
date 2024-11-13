package todos

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"test-todo/api/internal/domain/todos"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
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

	updateItemMatch := mock.MatchedBy(func(input *dynamodb.UpdateItemInput) bool {
		return input.TableName != nil &&
			*input.TableName == "testTable" &&
			input.Key != nil &&
			reflect.DeepEqual(input.Key, map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{
					Value: toDo.ID,
				},
			}) &&
			input.UpdateExpression != nil &&
			*input.UpdateExpression == "SET #t = :t, #c = :c, #ua = :ua" &&
			input.ExpressionAttributeNames != nil &&
			reflect.DeepEqual(input.ExpressionAttributeNames, map[string]string{
				"#t":  "title",
				"#c":  "completed",
				"#ua": "updated_at",
			}) &&
			input.ExpressionAttributeValues != nil &&
			reflect.DeepEqual(input.ExpressionAttributeValues, map[string]types.AttributeValue{
				":t": &types.AttributeValueMemberS{
					Value: toDo.Title,
				},
				":c": &types.AttributeValueMemberBOOL{
					Value: toDo.Completed,
				},
				":ua": &types.AttributeValueMemberS{
					Value: toDo.UpdatedAt.Format(time.RFC3339),
				},
			})
	})

	t.Run("successful update", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}

		mockDBClient.On("UpdateItem", contextMatch, updateItemMatch, optsMatch).
			Return(&dynamodb.UpdateItemOutput{}, nil)

		err := repo.Update(ctx, toDo)
		assert.NoError(t, err)
	})

	t.Run("update failure", func(t *testing.T) {
		mockDBClient := new(mockDynamoDBClient)
		repo := &repo{
			dbClient:  mockDBClient,
			tableName: "testTable",
		}
		mockDBClient.On("UpdateItem", contextMatch, updateItemMatch, optsMatch).
			Return(nil, errors.New("update error"))

		err := repo.Update(ctx, toDo)
		assert.Error(t, err)
		assert.Equal(t, "update error", err.Error())
	})
}
