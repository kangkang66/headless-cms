package cms

import "context"

type Field interface {
	GetName(ctx context.Context) (name string)
	GetNickName(ctx context.Context) (name string)
}

//基本结构
//字段
type BaseField struct {
	//xxxxxxx
	//名称
	Name string `json:"name"`
	//暴露给外部的名称
	NickName string `json:"nick_name"`
	//文字描述
	Desc string `json:"desc"`
	//状态 0删除 1正常
	Status int `json:"status"`

	//其他属性
	StorageType string      `json:"storage_type"`
	ShowType    string      `json:"show_type"`
	//引用的数据模型名称
	LinkModelName string `json:"link_model_name"`
	Default     interface{} `json:"default"`
	Valid       []ValidRule `json:"valid"`
	IsShow bool        `json:"is_show"`
	IsUnique    bool        `json:"is_unique"`
	IsMust      bool        `json:"is_must"`
}

func (b BaseField) GetName(ctx context.Context) (name string) {
	return b.Name
}
func (b BaseField) GetNickName(ctx context.Context) (name string) {
	return b.NickName
}

type ShortTextField struct {
	BaseField
}

func InitShortTextField() Field {
	return ShortTextField{BaseField{
		Name:        "",
		NickName:    "",
		Desc:        "",
		Status:      0,
		StorageType: "",
		ShowType:    "",
		Default:     nil,
		Valid:       nil,
		IsShow: false,
		IsUnique:    false,
		IsMust:      false,
	}}
}

type LongTextField struct {
	BaseField
}
type NumberField struct {
	BaseField
}
type EnumField struct {
	BaseField
}
type TimeField struct {
	BaseField
}
type UploadField struct {
	BaseField
}
type LinkField struct {
	BaseField
}
