package builder

import (
	"gitea.ipicture.vip/jerry/db-sdk/pbfiles"
	"google.golang.org/protobuf/types/known/structpb"
)

type ParamBuilder struct {
	param map[string]interface{}
}

func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{param: make(map[string]interface{})}
}

func (b *ParamBuilder) Add(name string, value interface{}) *ParamBuilder {
	b.param[name] = value
	return b
}

func (b *ParamBuilder) Build() *pbfiles.SimpleParams {
	paramStruct, _ := structpb.NewStruct(b.param)
	return &pbfiles.SimpleParams{
		Params: paramStruct,
	}
}
