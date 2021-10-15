package api

import (
	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/model"
)

func returnError(err error) *model.JSONResult {
	return &model.JSONResult{
		Code: 1,
		Err:  err.Error(),
	}
}

func returnData(value interface{}) *model.JSONResult {
	return &model.JSONResult{
		Code: 0,
		Data: value,
	}
}
