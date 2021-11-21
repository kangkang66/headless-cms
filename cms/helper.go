package cms

const (
	FieldStorageTypeVarchar	=	"varchar"
	FieldStorageTypeText	=	"text"

	FieldStorageTypeLinkModelOneOne	=	"link_model_one_one"
	FieldStorageTypeLinkModelOneMany	=	"link_model_one_many"
)

func IsLinkModel(storageType string) bool {
	var is bool
	if storageType == FieldStorageTypeLinkModelOneOne{
		is = true
	}else if storageType == FieldStorageTypeLinkModelOneMany {
		is = true
	}
	return is
}

func CloneMap(src map[string]struct{}) (new map[string]struct{}) {
	new = make(map[string]struct{})
	for k,v := range src {
		new[k] = v
	}
	return
}