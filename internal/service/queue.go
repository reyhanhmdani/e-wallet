package service

import (
	"e-wallet/domain"
	"e-wallet/internal/component"
	"e-wallet/internal/config"
	"github.com/hibiken/asynq"
)

type queueService struct {
	queueClient *asynq.Client
}

func NewQueue(cnf *config.Config) domain.QueueService {
	// harus memasang redis yang sama, kita memakai configurasi yang sama dengan e-wallet queue, supaya ponting nya sama
	redisConnection := asynq.RedisClientOpt{
		Addr:     cnf.Queue.Addr,
		Password: cnf.Queue.Password,
	}

	// initiate
	client := asynq.NewClient(redisConnection)

	return &queueService{
		queueClient: client,
	}
}

func (q queueService) Enqueue(name string, data []byte, retry int) error {
	task := asynq.NewTask(name, data, asynq.MaxRetry(retry))

	info, err := q.queueClient.Enqueue(task)
	if err != nil {
		component.Log.Error("error di bagian Service Enqueue", err.Error())
		return err
	}

	component.Log.Info("enqueue-client-id: ", info.ID)
	return nil
}
