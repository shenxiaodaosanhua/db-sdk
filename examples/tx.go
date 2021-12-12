package examples

import (
	"context"
	"gitea.ipicture.vip/jerry/db-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type UserTx struct {
	Id       int    `mapstructure:"id"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type ExecResult struct {
	RowsAffected int `mapstructure:"_RowsAffected"`
}

func TxTest() {
	client, err := builder.NewClientBuilder("localhost:8080").
		WithOption(grpc.WithInsecure()).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	tx := builder.NewTxBuilder(context.Background(), client)

	err = tx.Tx(func(tx *builder.TxBuilder) error {
		user := &User{
			Username: "jerry22",
			Mobile:   "18011801991",
			Password: "123123",
		}
		params := builder.NewParamBuilder().
			Add("username", user.Username).
			Add("mobile", user.Mobile).
			Add("password", user.Password)

		err := tx.Exec("add_user", params, user)
		if err != nil {
			return err
		}

		log.Println("新增用户ID：", user.Id)
		//log.Fatal("手动取消")
		params = builder.NewParamBuilder().Add("user_id", user.Id).Add("amount", 333)
		result := &ExecResult{}
		err = tx.Exec("add_user_amounts", params, result)
		if err != nil {
			return err
		}
		return nil
	})
	log.Println(err)
	select {}
}
