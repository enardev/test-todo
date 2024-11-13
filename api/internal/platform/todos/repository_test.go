package todos

import (
	"test-todo/api/cmd/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	cfg := config.DbConfig{
		Region:    "us-west-2",
		Endpoint:  "http://localhost:8000",
		AccessKey: "dummyAccessKey",
		SecretKey: "dummySecretKey",
		TableName: "test-table",
	}

	rep := NewRepository(cfg)

	assert.NotNil(t, rep)
	assert.Equal(t, cfg.TableName, rep.(*repo).tableName)
	assert.NotNil(t, rep.(*repo).dbClient)
}
