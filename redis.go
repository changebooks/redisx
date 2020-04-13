package redisx

import (
	"errors"
	"github.com/changebooks/log"
	"github.com/changebooks/redis"
)

type Redis struct {
	executor *redis.Executor
	schema   *redis.Schema
	logger   *log.Logger
}

func New(idRegister *log.IdRegister, logger *log.Logger, data map[string]string) (*Redis, error) {
	tag := "New"

	if logger == nil {
		return nil, errors.New("logger can't be nil")
	}

	if data == nil {
		err := errors.New("data can't be nil")
		logger.E(tag, Failure, nil, err, "", idRegister)
		return nil, err
	}

	profile, err := redis.NewProfile(data)
	if err != nil {
		logger.E(tag, Failure, data, err, "", idRegister)
		return nil, err
	}

	schema, err := redis.NewSchema(profile)
	if err != nil {
		logger.E(tag, Failure, profile, err, "", idRegister)
		return nil, err
	}

	executor, err := redis.NewExecutor(schema)
	if err != nil {
		logger.E(tag, Failure, schema, err, "", idRegister)
		return nil, err
	}

	return &Redis{
		executor: executor,
		schema:   schema,
		logger:   logger,
	}, nil
}

func (x *Redis) GetExecutor() *redis.Executor {
	return x.executor
}

func (x *Redis) GetSchema() *redis.Schema {
	return x.schema
}

func (x *Redis) GetLogger() *log.Logger {
	return x.logger
}
