package redisx

import (
	"github.com/changebooks/log"
	"time"
)

func (x *Redis) HGet(idRegister *log.IdRegister, key string, field string) (value interface{}, err error) {
	tag := "HGet"

	start := time.Now()

	value, err, closeErr := x.executor.HGet(key, field)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HGET", nil, key, field, nil, value, 0)

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

func (x *Redis) HSet(idRegister *log.IdRegister, key string, field string, value interface{}) (result interface{}, err error) {
	tag := "HSet"

	start := time.Now()

	result, err, closeErr := x.executor.HSet(key, field, value)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HSET", nil, key, field, value, result, 0)

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

func (x *Redis) HDel(idRegister *log.IdRegister, key string, fields ...string) (affectedRows interface{}, err error) {
	tag := "HDel"

	start := time.Now()

	affectedRows, err, closeErr := x.executor.HDel(key, fields...)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HDEL", nil, key, fields, nil, affectedRows, 0)

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

func (x *Redis) HExists(idRegister *log.IdRegister, key string, field string) (result interface{}, err error) {
	tag := "HExists"

	start := time.Now()

	result, err, closeErr := x.executor.HExists(key, field)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HEXISTS", nil, key, field, nil, result, 0)

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

func (x *Redis) HMGet(idRegister *log.IdRegister, key string, fields ...string) (values []interface{}, err error) {
	tag := "HMGet"

	start := time.Now()

	values, err, closeErr := x.executor.HMGet(key, fields...)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HMGET", nil, key, fields, nil, values, 0)

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

func (x *Redis) HMSet(idRegister *log.IdRegister, key string, data map[string]interface{}) (result interface{}, err error) {
	tag := "HMSet"

	start := time.Now()

	result, err, closeErr := x.executor.HMSet(key, data)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HMSET", nil, key, "", data, result, 0)

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

func (x *Redis) HGetAll(idRegister *log.IdRegister, key string) (values []interface{}, err error) {
	tag := "HGetAll"

	start := time.Now()

	values, err, closeErr := x.executor.HGetAll(key)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HGETALL", nil, key, "", nil, values, 0)

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

func (x *Redis) HKeys(idRegister *log.IdRegister, key string) (fields []string, err error) {
	tag := "HKeys"

	start := time.Now()

	values, err, closeErr := x.executor.HKeys(key)
	if values != nil {
		for v := range values {
			fields = append(fields, string(v))
		}
	}

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HKEYS", nil, key, "", nil, values, 0)

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

func (x *Redis) HVals(idRegister *log.IdRegister, key string) (values []interface{}, err error) {
	tag := "HVals"

	start := time.Now()

	values, err, closeErr := x.executor.HVals(key)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "HVALS", nil, key, "", nil, values, 0)

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
