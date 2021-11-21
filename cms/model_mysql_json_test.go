package cms

import (
	"context"
	"fmt"
	"testing"
)

func TestModelMysqlJson_Create_Source(t *testing.T) {
	c := newModelMysqlJson()
	ctx := context.Background()
	mod := &modelElement{
		Name:   "source",
		Desc:   "购买渠道",
		Status: 0,
		Fields: []BaseField{
			{
				Name:          "name",
				NickName:      "name",
				Desc:          "渠道名称",
				Status:        0,
				StorageType:   "varchar",
				ShowType:      "string",
				LinkModelName: "",
				Default:       "",
				Valid:         nil,
				IsShow:        true,
				IsUnique:      false,
				IsMust:        false,
			},
		},
	}
	id,err := c.Create(ctx,mod.Name,mod)
	fmt.Println(id,err)
}

func TestModelMysqlJson_Create_Like(t *testing.T) {
	c := newModelMysqlJson()
	ctx := context.Background()
	mod := &modelElement{
		Name:   "like",
		Desc:   "喜欢",
		Status: 0,
		Fields: []BaseField{
			{
				Name:          "name",
				NickName:      "name",
				Desc:          "名称",
				Status:        0,
				StorageType:   "varchar",
				ShowType:      "string",
				LinkModelName: "",
				Default:       "",
				Valid:         nil,
				IsShow:        true,
				IsUnique:      false,
				IsMust:        false,
			},
			{
				Name:          "buy_source",
				NickName:      "buy_source",
				Desc:          "购买渠道",
				Status:        0,
				StorageType:   "link_model_one_one",
				ShowType:      "link_model_one_one",
				LinkModelName: "source",
				Default:       "",
				Valid:         nil,
				IsShow:        true,
				IsUnique:      false,
				IsMust:        false,
			},
		},
	}
	id,err := c.Create(ctx,mod.Name,mod)
	fmt.Println(id,err)
}

func TestModelMysqlJson_Create_LinkRoundErr(t *testing.T) {
	c := newModelMysqlJson()
	ctx := context.Background()
	mod := &modelElement{
		Name:   "user",
		Desc:   "用户",
		Status: 0,
		Fields: []BaseField{
			{
				Name:          "name",
				NickName:      "name",
				Desc:          "姓名",
				Status:        0,
				StorageType:   "varchar",
				ShowType:      "string",
				LinkModelName: "",
				Default:       "",
				Valid:         nil,
				IsShow:        true,
				IsUnique:      false,
				IsMust:        false,
			},
			{
				Name:          "likes",
				NickName:      "likes",
				Desc:          "爱好",
				Status:        0,
				StorageType:   "link_model_one_many",
				ShowType:      "link_model_one_many",
				LinkModelName: "like",
				Default:       "",
				Valid:         nil,
				IsShow:        true,
				IsUnique:      false,
				IsMust:        false,
			},
			{
				Name:          "user_source",
				NickName:      "user_source",
				Desc:          "渠道",
				Status:        0,
				StorageType:   "link_model_one_one",
				ShowType:      "link_model_one_one",
				LinkModelName: "user",
				Default:       "",
				Valid:         nil,
				IsShow:        true,
				IsUnique:      false,
				IsMust:        false,
			},
		},
	}
	id,err := c.Create(ctx,mod.Name,mod)
	fmt.Println(id,err)
}


func TestModelMysqlJson_FindOne(t *testing.T) {
	c := newModelMysqlJson()
	ctx := context.Background()

	mod,err := c.FindOne(ctx,"user")
	fmt.Println(mod,err)
}
