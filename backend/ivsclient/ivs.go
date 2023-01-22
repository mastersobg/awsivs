package ivsclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ivs"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
)

type Client struct {
	client *ivs.Client
}

func NewClient(ctx context.Context) (*Client, error) {
	// Load the shared AWS configuration from ~/.aws/config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	// Create an IVS service client
	return &Client{
		client: ivs.NewFromConfig(cfg),
	}, nil
}

func (c *Client) CreateChannel(ctx context.Context, name string) (*types.Channel, *types.StreamKey, error) {
	resp, err := c.client.CreateChannel(ctx, &ivs.CreateChannelInput{
		Name: aws.String(name),
	})
	if err != nil {
		return nil, nil, err
	}
	return resp.Channel, resp.StreamKey, nil
}

func (c *Client) GetStream(ctx context.Context, channelARN *string) (*types.Stream, error) {
	resp, err := c.client.GetStream(ctx, &ivs.GetStreamInput{
		ChannelArn: channelARN,
	})
	if err != nil {
		return nil, err
	}
	return resp.Stream, nil
}

func (c *Client) PutMetadata(ctx context.Context, channelARN *string, metadata string) error {
	_, err := c.client.PutMetadata(ctx, &ivs.PutMetadataInput{
		ChannelArn: channelARN,
		Metadata:   &metadata,
	})
	return err
}

func (c *Client) DeleteChannel(ctx context.Context, channelARN *string) error {
	_, err := c.client.DeleteChannel(ctx, &ivs.DeleteChannelInput{
		Arn: channelARN,
	})
	return err
}
