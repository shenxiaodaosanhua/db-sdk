package builder

import (
	"context"
	"gitea.ipicture.vip/jerry/db-sdk/pbfiles"
	"google.golang.org/grpc"
)

type ClientBuilder struct {
	url  string
	opts []grpc.DialOption
}

func NewClientBuilder(url string) *ClientBuilder {
	return &ClientBuilder{url: url}
}

func (b *ClientBuilder) WithOption(opts ...grpc.DialOption) *ClientBuilder {
	b.opts = append(b.opts, opts...)
	return b
}

func (b *ClientBuilder) Build() (pbfiles.DBServiceClient, error) {
	cc, err := grpc.DialContext(context.Background(), b.url, b.opts...)
	if err != nil {
		return nil, err
	}

	return pbfiles.NewDBServiceClient(cc), nil
}
