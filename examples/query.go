package examples

import (
	"context"
	"fmt"
	"gitea.ipicture.vip/jerry/db-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
)

type User struct {
	Id        int    `mapstructure:"id"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Mobile    string `mapstructure:"mobile"`
	CreatedAt string `mapstructure:"created_at"`
	UpdatedAt string `mapstructure:"updated_at"`
}

func Query() {
	client, err := builder.NewClientBuilder("localhost:8080").
		WithOption(grpc.WithInsecure()).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	params := builder.NewParamBuilder().
		Add("id", 1)

	api := builder.NewApiBuilder("userlist", builder.APITYPE_QUERY)
	result := make([]*User, 0)
	err = api.Invoke(context.Background(), params, client, &result)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range result {
		fmt.Println(user)
	}
}
