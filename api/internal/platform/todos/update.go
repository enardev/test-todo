package todos

import (
	"context"
	"test-todo/api/internal/domain/todos"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (r *repo) Update(ctx context.Context, todo todos.ToDo) error {
	toUpdate := mapToEntity(todo)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: toUpdate.ID,
			},
		},
		UpdateExpression: aws.String("SET #t = :t, #c = :c, #ua = :ua"),
		ExpressionAttributeNames: map[string]string{
			"#t":  "title",
			"#c":  "completed",
			"#ua": "updated_at",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":t": &types.AttributeValueMemberS{
				Value: toUpdate.Title,
			},
			":c": &types.AttributeValueMemberBOOL{
				Value: toUpdate.Completed,
			},
			":ua": &types.AttributeValueMemberS{
				Value: toUpdate.UpdatedAt.Format(time.RFC3339),
			},
		},
	}

	_, err := r.dbClient.UpdateItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
