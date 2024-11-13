package todos

import (
	"context"
	"log"
	"test-todo/api/cmd/config"
	"test-todo/api/internal/domain/todos"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Repository interface {
	FindAll(context.Context) ([]ToDo, error)
	FindByID(context.Context, string) (ToDo, error)
	Exists(context.Context, string) (bool, error)
	Save(context.Context, ToDo) error
	Update(context.Context, ToDo) error
	Delete(context.Context, string) error
}

type DynamoDbAPI interface {
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
}

type repo struct {
	tableName string
	dbClient  DynamoDbAPI
}

func NewRepository(cfg config.DbConfig) todos.Repository {
	awsCfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		awsconfig.WithRegion(cfg.Region),
		awsconfig.WithBaseEndpoint(cfg.Endpoint),
		awsconfig.WithCredentialsProvider(
			aws.NewCredentialsCache(
				credentials.NewStaticCredentialsProvider(
					cfg.AccessKey,
					cfg.SecretKey,
					"",
				),
			),
		),
	)

	if err != nil {
		log.Fatalf("failed to load AWS config: %v\n", err)
	}

	client := dynamodb.NewFromConfig(awsCfg)

	return &repo{
		tableName: cfg.TableName,
		dbClient:  client,
	}
}
