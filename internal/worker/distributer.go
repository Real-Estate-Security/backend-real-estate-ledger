package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

// allows to mock for testing
type TaskDistributer interface {
	// DistributeTask distributes the task to the worker.
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributer struct {
	// redisClient is the redis client.
	client *asynq.Client
}

func NewRedisTaskDistributer(redisOpt asynq.RedisClientOpt) TaskDistributer {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributer{client: client}
}
