package redisx

import (
	"github.com/changebooks/log"
	"time"
)

func (x *Redis) HIncrBy(idRegister *log.IdRegister, key string, field string, increment int) (value interface{}, err error) {
	tag := "HIncrBy"

	start := time.Now()

	value, err, closeErr := x.executor.HIncrBy(key, field, increment)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HINCRBY", nil, key, field, increment, value, 0)

	if closeErr != nil {
		x.logger.E(tag, CloseFailure, remark, closeErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, Success, remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}

func (x *Redis) HIncrByFloat(idRegister *log.IdRegister, key string, field string, increment float64) (value interface{}, err error) {
	tag := "HIncrByFloat"

	start := time.Now()

	value, err, closeErr := x.executor.HIncrByFloat(key, field, increment)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HINCRBYFLOAT", nil, key, field, increment, value, 0)

	if closeErr != nil {
		x.logger.E(tag, CloseFailure, remark, closeErr, "", idRegister)
	}

	if err == nil {
		x.logger.I(tag, Success, remark, idRegister)
	} else {
		x.logger.E(tag, Failure, remark, err, "", idRegister)
	}

	return
}
