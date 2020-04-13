package redisx

import (
	"github.com/changebooks/log"
	"time"
)

func (x *Redis) Incr(idRegister *log.IdRegister, key string) (value interface{}, err error) {
	tag := "Incr"

	start := time.Now()

	value, err, closeErr := x.executor.Incr(key)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "INCR", nil, key, "", nil, value, 0)

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

func (x *Redis) Decr(idRegister *log.IdRegister, key string) (value interface{}, err error) {
	tag := "Decr"

	start := time.Now()

	value, err, closeErr := x.executor.Decr(key)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "DECR", nil, key, "", nil, value, 0)

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

func (x *Redis) IncrBy(idRegister *log.IdRegister, key string, increment int) (value interface{}, err error) {
	tag := "IncrBy"

	start := time.Now()

	value, err, closeErr := x.executor.IncrBy(key, increment)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "INCRBY", nil, key, "", increment, value, 0)

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

func (x *Redis) DecrBy(idRegister *log.IdRegister, key string, decrement int) (value interface{}, err error) {
	tag := "DecrBy"

	start := time.Now()

	value, err, closeErr := x.executor.DecrBy(key, decrement)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "DECRBY", nil, key, "", decrement, value, 0)

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

func (x *Redis) IncrByFloat(idRegister *log.IdRegister, key string, increment float64) (value interface{}, err error) {
	tag := "IncrByFloat"

	start := time.Now()

	value, err, closeErr := x.executor.IncrByFloat(key, increment)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "INCRBYFLOAT", nil, key, "", increment, value, 0)

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
