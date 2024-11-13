package todos

import (
	"context"
	"test-todo/api/internal/domain/todos"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *repo) Save(ctx context.Context, toDo todos.ToDo) error {
	toSave := mapToEntity(toDo)

	item, err := attributevalue.MarshalMap(toSave)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	}

	_, err = r.dbClient.PutItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
