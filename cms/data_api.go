package cms

import (
	"context"
)

type PageParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type DataApi interface {
	//数据操作接口
	//搜索
	FindOne(ctx context.Context, modelName string, id int64) (resp interface{}, err error)
	FindMany(ctx context.Context, modelName string, ids []int64) (resp interface{},err error)
	FindAll(ctx context.Context, modelName string, page PageParams, params map[string]interface{}) (resp interface{}, err error)
	Count(ctx context.Context, modelName string, params map[string]interface{}) (total int64, err error)
	//插入
	Insert(ctx context.Context, modelName string, params map[string]interface{}) (lastId uint, err error)
	//更新
	Update(ctx context.Context, modelName string, id int64, fields map[string]interface{}) (rowsAffected int64, err error)
	//删除
	Delete(ctx context.Context, modelName string, id int64) (rowsAffected int64, err error)
}

func NewDataApi() DataApi {
	return newDataMysqlJson()
}