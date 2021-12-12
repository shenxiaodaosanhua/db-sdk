package builder

import (
	"context"
	"gitea.ipicture.vip/jerry/db-sdk/pbfiles"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
	"log"
)

const (
	TX_EXEC  = "exec"
	TX_QUERY = "query"
)

type TxBuilder struct {
	ctx    context.Context
	cancel context.CancelFunc
	client pbfiles.DBService_TxClient
}

func NewTxBuilder(
	ctx context.Context,
	client pbfiles.DBServiceClient,
	opts ...grpc.CallOption) *TxBuilder {
	txCtx, cancel := context.WithCancel(ctx)
	txClient, err := client.Tx(txCtx, opts...)
	if err != nil {
		panic(err)
	}

	return &TxBuilder{
		ctx:    txCtx,
		client: txClient,
		cancel: cancel,
	}
}

func (b *TxBuilder) Exec(name string, params *ParamBuilder, out interface{}) error {
	err := b.client.Send(&pbfiles.TxRequest{
		Name:   name,
		Params: params.Build(),
		Type:   TX_EXEC,
	})
	if err != nil {
		return err
	}

	rsp, err := b.client.Recv()
	if err != nil {
		return err
	}

	if out == nil {
		return nil
	}

	if execRet, ok := rsp.Result.AsMap()["exec"]; ok {
		if execRet.([]interface{})[1] != nil {
			m := execRet.([]interface{})[1].(map[string]interface{})
			return mapstructure.Decode(m, out)
		}

		m := map[string]interface{}{
			"_RowsAffected": execRet.([]interface{})[0],
		}
		return mapstructure.Decode(m, out)
	}

	return nil
}

func (b *TxBuilder) Query(name string, params *ParamBuilder, out interface{}) error {
	err := b.client.Send(&pbfiles.TxRequest{
		Name:   name,
		Type:   TX_QUERY,
		Params: params.Build(),
	})
	if err != nil {
		return err
	}

	rsp, err := b.client.Recv()
	if err != nil {
		return err
	}

	if out == nil {
		return nil
	}

	if queryRet, ok := rsp.Result.AsMap()["query"]; ok {
		return mapstructure.Decode(queryRet, out)
	}
	return nil
}

func (b *TxBuilder) Tx(fn func(tx *TxBuilder) error) error {
	err := fn(b)
	if err != nil {
		log.Println("tx error:", err)
		b.cancel()
		return err
	}

	return b.client.CloseSend()
}
