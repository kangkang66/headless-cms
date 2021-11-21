package cms

import (
	"context"
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DataMysqlJson struct {
	db *gorm.DB
}

func newDataMysqlJson() *DataMysqlJson {
	dsn := "root:strapi@tcp(127.0.0.1:3306)/cms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	api := &DataMysqlJson{
		db: db,
	}
	return api
}

type tableDatas struct {
	gorm.Model
	ModelName string `json:"model_name"`
	Data      []byte `json:"data"`
	Status	  int `json:"status"`
}
func (tableDatas) TableName() string {
	return "datas"
}

//搜索
func (d *DataMysqlJson) FindOne(ctx context.Context, modelName string, id int64) (resp interface{}, err error) {
	//获取数据

	//获取模型结构
	//分析模型字段是否有引用
	//获取引用对象的数据

	return
}

func (d *DataMysqlJson) FindMany(ctx context.Context, modelName string, ids []int64) (resp interface{},err error) {
	return
}
func (d *DataMysqlJson) FindAll(ctx context.Context, modelName string, page PageParams, params map[string]interface{}) (resp interface{}, err error) {
	return
}
func (d *DataMysqlJson) Count(ctx context.Context, modelName string, params map[string]interface{}) (total int64, err error) {
	return
}
//插入
func (d *DataMysqlJson) Insert(ctx context.Context, modelName string, params map[string]interface{}) (lastId uint, err error) {
	//todo 检查字段是否存在
	//检验值的合法性
	//插入数据
	jb,err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return
	}
	data := &tableDatas{
		ModelName: modelName,
		Data:      jb,
	}
	err = d.db.Create(data).Error
	if err != nil {
		log.Println(err)
		return
	}
	return data.ID,nil
}
//更新
func (d *DataMysqlJson) Update(ctx context.Context, modelName string, id int64, fields map[string]interface{}) (rowsAffected int64, err error) {
	return
}
//删除
func (d *DataMysqlJson) Delete(ctx context.Context, modelName string, id int64) (rowsAffected int64, err error) {
	return
}