package redisx

import (
	"github.com/changebooks/log"
	"time"
)

func (x *Redis) Get(idRegister *log.IdRegister, key string) (value interface{}, err error) {
	tag := "Get"

	start := time.Now()

	value, err, closeErr := x.executor.Get(key)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "GET", nil, key, "", nil, value, 0)

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

func (x *Redis) Set(idRegister *log.IdRegister, key string, value interface{}) (result interface{}, err error) {
	tag := "Set"

	start := time.Now()

	result, err, closeErr := x.executor.Set(key, value)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "SET", nil, key, "", value, result, 0)

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

func (x *Redis) SetEx(idRegister *log.IdRegister, key string, value interface{}, seconds int) (result interface{}, err error) {
	tag := "SetEx"

	start := time.Now()

	result, err, closeErr := x.executor.SetEx(key, value, seconds)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "SETEX", nil, key, "", value, result, seconds)

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

func (x *Redis) Del(idRegister *log.IdRegister, keys ...string) (affectedRows interface{}, err error) {
	tag := "Del"

	start := time.Now()

	affectedRows, err, closeErr := x.executor.Del(keys...)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "DEL", nil, keys, "", nil, affectedRows, 0)

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

func (x *Redis) Exists(idRegister *log.IdRegister, key string) (result interface{}, err error) {
	tag := "Exists"

	start := time.Now()

	result, err, closeErr := x.executor.Exists(key)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "EXISTS", nil, key, "", nil, result, 0)

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

// -1 : Forever
// -2 : Nothing
func (x *Redis) Ttl(idRegister *log.IdRegister, key string) (seconds interface{}, err error) {
	tag := "Ttl"

	start := time.Now()

	seconds, err, closeErr := x.executor.Ttl(key)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "TTL", nil, key, "", nil, seconds, 0)

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

func (x *Redis) MGet(idRegister *log.IdRegister, keys ...string) (values []interface{}, err error) {
	tag := "MGet"

	start := time.Now()

	values, err, closeErr := x.executor.MGet(keys...)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "MGET", nil, keys, "", nil, values, 0)

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

func (x *Redis) MSet(idRegister *log.IdRegister, data map[string]interface{}) (result interface{}, err error) {
	tag := "MSet"

	start := time.Now()

	result, err, closeErr := x.executor.MSet(data)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "MSET", nil, "", "", data, result, 0)

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
