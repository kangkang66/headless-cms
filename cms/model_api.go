package cms

import (
	"context"
	"gorm.io/gorm"
)

type modelElement struct {
	gorm.Model
	Name   string      `json:"name"`
	Desc   string      `json:"desc"`
	Status int         `json:"status"`
	//存储方式
	Fields []BaseField `json:"fields"`
}

type ModelApi interface {
	//模型操作接口
	//创建
	Create(ctx context.Context, modelName string, modelStruct *modelElement) (lastId uint, err error)
	//搜索
	FindOne(ctx context.Context, modelName string) (resp *modelElement, err error)
	//更新
	//新增字段
	UpdateFieldAdd(ctx context.Context, modelName string, field BaseField) (rowsAffected int64, err error)
	//删除字段
	UpdateFieldDel(ctx context.Context, modelName string, fieldName string) (rowsAffected int64, err error)

	//删除
	Delete(ctx context.Context, modelName string) (rowsAffected int64, err error)
}

func NewModelApi() ModelApi {
	return newModelMysqlJson()
}
