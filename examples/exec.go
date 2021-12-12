package examples

import (
	"context"
	"fmt"
	"gitea.ipicture.vip/jerry/db-sdk/pbfiles"
	"gitea.ipicture.vip/jerry/db-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type UserAddResult struct {
	UserId       int `mapstructure:"user_id"`
	RowsAffected int `mapstructure:"user_id"`
}

func Exec() {
	cc, err := grpc.DialContext(context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	client := pbfiles.NewDBServiceClient(cc)

	params := builder.NewParamBuilder().
		Add("username", "jialiang7").
		Add("mobile", "18011801988").
		Add("password", "123123")

	api := builder.NewApiBuilder("add_user", builder.APITYPE_EXEC)
	result := &UserAddResult{}
	err = api.Invoke(context.Background(), params, client, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
