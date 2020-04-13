package redisx

import (
	"github.com/changebooks/redis"
	"time"
)

type Remark struct {
	Proto       string        `json:"proto"`
	Database    int           `json:"database"`
	Address     string        `json:"address"`
	CommandName string        `json:"commandName"`
	Args        interface{}   `json:"args"`
	Key         interface{}   `json:"key"`
	Field       interface{}   `json:"field"`
	Value       interface{}   `json:"value"`
	Result      interface{}   `json:"result"`
	Seconds     int           `json:"seconds"`
	Total       time.Duration `json:"total"`
	Start       time.Time     `json:"start"`
	Done        time.Time     `json:"done"`
}

func NewRemark(schema *redis.Schema, start time.Time, done time.Time, commandName string, args interface{},
	key interface{}, field interface{}, value interface{}, result interface{}, seconds int) *Remark {
	proto := ""
	database := DatabaseUnknown
	address := ""
	if schema != nil {
		proto = schema.GetProto()
		database = schema.GetDatabase()
		address = schema.GetAddress()
	}

	total := done.Sub(start)

	return &Remark{
		Proto:       proto,
		Database:    database,
		Address:     address,
		CommandName: commandName,
		Args:        args,
		Key:         key,
		Field:       field,
		Value:       value,
		Result:      result,
		Seconds:     seconds,
		Total:       total,
		Start:       start,
		Done:        done,
	}
}
