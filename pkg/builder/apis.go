package builder

import (
	"context"
	"gitea.ipicture.vip/jerry/db-sdk/helpers"
	"gitea.ipicture.vip/jerry/db-sdk/pbfiles"
	"github.com/mitchellh/mapstructure"
)

const (
	APITYPE_QUERY = iota
	APITYPE_EXEC
	APITYPE_TX
)

type ApiBuilder struct {
	name    string
	apiType int
}

func NewApiBuilder(name string, apiType int) *ApiBuilder {
	return &ApiBuilder{name: name, apiType: apiType}
}

func (b *ApiBuilder) Invoke(
	ctx context.Context,
	params *ParamBuilder,
	client pbfiles.DBServiceClient, out interface{}) error {
	if b.apiType == APITYPE_QUERY {
		request := &pbfiles.QueryRequest{
			Name:   b.name,
			Params: params.Build(),
		}

		response, err := client.Query(ctx, request)
		if err != nil {
			return err
		}
		list := helpers.PbStructsToMapList(response.GetResult())
		return mapstructure.Decode(list, out)
	}

	if b.apiType == APITYPE_EXEC {
		request := &pbfiles.ExecRequest{
			Name:   b.name,
			Params: params.Build(),
		}

		response, err := client.Exec(ctx, request)
		if err != nil {
			return err
		}

		var m map[string]interface{}
		if response.Select != nil {
			m = response.Select.AsMap()
			m["_RowsAffected"] = response.RowsAffected
		} else {
			m = map[string]interface{}{
				"_RowsAffected": response.RowsAffected,
			}
		}
		return mapstructure.Decode(m, out)
	}

	return nil
}
