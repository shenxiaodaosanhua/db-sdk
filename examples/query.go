package examples

import (
	"context"
	"fmt"
	"gitea.ipicture.vip/jerry/db-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
	"time"
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

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	defer cancel()
	err = api.Invoke(ctx, params, client, &result)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range result {
		fmt.Println(user)
	}
}
