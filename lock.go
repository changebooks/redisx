package redisx

import (
	"github.com/changebooks/log"
	"time"
)

func (x *Redis) Lock(idRegister *log.IdRegister, name string, token string, seconds int) (result interface{}, err error) {
	tag := "Lock"

	start := time.Now()

	result, err, closeErr := x.executor.Lock(name, token, seconds)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "SET-EX-NX", nil, name, "", token, result, seconds)

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

func (x *Redis) Unlock(idRegister *log.IdRegister, name string, token string) (result interface{}, err error) {
	tag := "Unlock"

	start := time.Now()

	result, err, closeErr := x.executor.Unlock(name, token)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "SCRIPT-DEL", nil, name, "", token, result, 0)

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
