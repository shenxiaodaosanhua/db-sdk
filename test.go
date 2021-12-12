package main

import (
	"context"
	"fmt"
	"gitea.ipicture.vip/jerry/db-sdk/pbfiles"
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

func main() {
	cc, err := grpc.DialContext(context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	client := pbfiles.NewDBServiceClient(cc)

	paramBuilder := builder.NewParamBuilder().Add("id", 1)

	users := make([]*User, 0)
	err = builder.NewApiBuilder("userlist", builder.APITYPE_QUERY).
		Invoke(context.Background(), paramBuilder, client, &users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
