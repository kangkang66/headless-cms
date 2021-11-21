package cms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func init()  {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type ModelMysqlJson struct {
	db *gorm.DB
}

var (
	ErrModelLinkRound = errors.New("model link round")
)

type tableModels struct {
	gorm.Model
	Name   string      `json:"name"`
	Desc   string      `json:"desc"`
	Status int         `json:"status"`
	Fields []byte 		`json:"fields"`
}
func (tableModels) TableName() string {
	return "models"
}

func newModelMysqlJson() ModelApi {
	dsn := "root:strapi@tcp(127.0.0.1:3306)/cms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mod := &ModelMysqlJson{
		db: db,
	}
	return mod
}

//检查组件引用关系，别有死循环
func (m *ModelMysqlJson) Create(ctx context.Context, modelName string, modelStruct *modelElement) (lastId uint, err error) {
	usedModStorage := map[string]struct{}{
		modelName: {},
	}
	err = m.checkFieldlinkRound(ctx, modelStruct.Fields, usedModStorage)
	if err != nil {
		return
	}

	fieldJSON,err := json.Marshal(modelStruct.Fields)
	if err != nil {
		log.Println(err)
		return
	}
	insert := &tableModels{
		Name:   modelName,
		Desc:   modelStruct.Desc,
		Fields: fieldJSON,
	}
	err = m.db.Create(insert).Error
	if err != nil {
		log.Println(err)
		return
	}
	return insert.ID,nil
}

func (m *ModelMysqlJson) checkFieldlinkRound(ctx context.Context, fields []BaseField, usedModStorage map[string]struct{}) (err error) {

	for _,field := range fields {
		if IsLinkModel(field.StorageType) {
			//多个字段key引用同一个mod，比如modA的a,b字段可以引用modB
			tmpStorage := CloneMap(usedModStorage)
			fmt.Println(field.Name, tmpStorage)

			//引用了已经出现的模型，死循环
			if _,exist := tmpStorage[field.LinkModelName];exist {
				err = ErrModelLinkRound
				log.Println(field.LinkModelName , tmpStorage, err)
				return
			}
			//获取这个link model的结构体
			var linkModStruct *modelElement
			linkModStruct,err = m.FindOne(ctx, field.LinkModelName)
			if err != nil {
				return
			}
			tmpStorage[field.LinkModelName] = struct{}{}

			err = m.checkFieldlinkRound(ctx, linkModStruct.Fields, tmpStorage)
			if err != nil {
				return
			}
		}
	}

	return
}


//搜索
func (m *ModelMysqlJson) FindOne(ctx context.Context, modelName string) (modelStruct *modelElement, err error) {
	modData := new(tableModels)
	err = m.db.Model(modData).Where("name=?",modelName).Find(modData).Error
	if modData.ID == 0 {
		err = gorm.ErrRecordNotFound
	}
	if err != nil {
		log.Println(err)
		return
	}


	fields := []BaseField{}
	err = json.Unmarshal(modData.Fields, &fields)
	if err != nil {
		log.Println(err)
		return
	}
	modelStruct = &modelElement{
		Model: modData.Model,
		Name:   modData.Name,
		Desc:   modData.Desc,
		Status: modData.Status,
		Fields: fields,
	}
	return
}

//更新
//新增字段
func (m *ModelMysqlJson) UpdateFieldAdd(ctx context.Context, modelName string, field BaseField) (rowsAffected int64, err error) {
	//检查有没有出现重复
	//如果是link类型，检查有没有死循环
	//追加到fields中

	return
}

//删除字段
func (m *ModelMysqlJson) UpdateFieldDel(ctx context.Context, modelName string, fieldName string) (rowsAffected int64, err error) {
	//field status = 1 表示删除

	return
}

//删除
func (m *ModelMysqlJson) Delete(ctx context.Context, modelName string) (rowsAffected int64, err error) {
	//检查有没有数据，如果有数据禁止删除
	//把模型 status = 1
	return
}
