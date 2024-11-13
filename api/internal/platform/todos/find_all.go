package todos

import (
	"context"
	"test-todo/api/internal/domain/todos"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r *repo) FindAll(ctx context.Context) ([]todos.ToDo, error) {
	result, err := r.dbClient.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	})
	if err != nil {
		return []todos.ToDo{}, err
	}

	var toDos []ToDo
	err = attributevalue.UnmarshalListOfMaps(result.Items, &toDos)
	if err != nil {
		return []todos.ToDo{}, err
	}

	return mapToDomainList(toDos), nil
}
