package cms

import "context"

type DataMongo struct {

}

//搜索
func (d *DataMongo) FindOne(ctx context.Context, modelName string, id int64) (resp interface{}, err error) {
	return
}
func (d *DataMongo) FindMany(ctx context.Context, modelName string, ids []int64) (resp interface{},err error) {
	return
}

func (d *DataMongo) FindAll(ctx context.Context, modelName string, page PageParams, params map[string]interface{}) (resp interface{}, err error) {
	return
}
func (d *DataMongo) Count(ctx context.Context, modelName string, params map[string]interface{}) (total int64, err error) {
	return
}
//插入
func (d *DataMongo) Insert(ctx context.Context, modelName string, params map[string]interface{}) (lastId uint, err error) {
	return
}
//更新
func (d *DataMongo) Update(ctx context.Context, modelName string, id int64, fields map[string]interface{}) (rowsAffected int64, err error) {
	return
}
//删除
func (d *DataMongo) Delete(ctx context.Context, modelName string, id int64) (rowsAffected int64, err error) {
	return
}