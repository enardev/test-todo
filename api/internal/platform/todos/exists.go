package todos

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (r *repo) Exists(ctx context.Context, id string) (bool, error) {
	input := dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	}

	result, err := r.dbClient.GetItem(ctx, &input)
	if err != nil {
		return false, err
	}

	if result.Item != nil {
		return true, nil
	}

	return false, nil
}
