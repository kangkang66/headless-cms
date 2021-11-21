package cms

import "context"

type ModelMongo struct {
}

//检查组件引用关系，别有死循环
func (ModelMongo) Create(ctx context.Context, modelName string, modelStruct *modelElement) (lastId uint, err error) {
	return
}

//搜索
func (ModelMongo) FindOne(ctx context.Context, modelName string) (resp *modelElement, err error) {
	return
}

//更新
//新增字段
func (ModelMongo) UpdateFieldAdd(ctx context.Context, modelName string, field BaseField) (rowsAffected int64, err error) {
	return
}

//删除字段
func (ModelMongo) UpdateFieldDel(ctx context.Context, modelName string, fieldName string) (rowsAffected int64, err error) {
	return
}

//删除
func (ModelMongo) Delete(ctx context.Context, modelName string) (rowsAffected int64, err error) {
	return
}
