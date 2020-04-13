package redisx

import (
	"github.com/changebooks/log"
	"github.com/gomodule/redigo/redis"
	"time"
)

func (x *Redis) Execute(idRegister *log.IdRegister, commandName string, args ...interface{}) (reply interface{}, err error) {
	tag := "Execute"

	start := time.Now()

	reply, err, closeErr := x.executor.Execute(commandName, args...)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, commandName, args, "", "", nil, reply, 0)

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

func (x *Redis) Script(idRegister *log.IdRegister, script *redis.Script, args ...interface{}) (reply interface{}, err error) {
	tag := "Script"

	start := time.Now()

	reply, err, closeErr := x.executor.Script(script, args...)

	done := time.Now()
	remark := NewRemark(x.schema, start, done, "script", args, "", "", nil, reply, 0)

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
