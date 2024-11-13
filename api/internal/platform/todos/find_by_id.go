package todos

import (
	"context"
	"errors"
	"test-todo/api/internal/domain/todos"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (r *repo) FindByID(ctx context.Context, id string) (todos.ToDo, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	}

	result, err := r.dbClient.GetItem(ctx, input)
	if err != nil {
		return todos.ToDo{}, err
	}

	if result.Item == nil {
		return todos.ToDo{}, errors.New("item not found")
	}

	var toDo ToDo
	err = attributevalue.UnmarshalMap(result.Item, &toDo)
	if err != nil {
		return todos.ToDo{}, err
	}

	return mapToDomain(toDo), nil
}
