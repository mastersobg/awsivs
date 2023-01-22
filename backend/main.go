package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/mastersobg/aws-ivs/ivsclient"
)

var (
	operation  *string
	channelArn *string
)

func initFlags() {
	operation = flag.String("operation", "", "--operation=[createChannel|deleteChannel|streamOnlineViewers]")
	channelArn = flag.String("channelArn", "", "--channelArn=<CHANNEL_ARN>")
	flag.Parse()
}

func streamOnlineViewers(ctx context.Context, client *ivsclient.Client) error {
	if channelArn == nil || len(*channelArn) == 0 {
		return errors.New("empty channel arn")
	}
	for {
		stream, err := client.GetStream(ctx, channelArn)
		if err != nil {
			return err
		}
		if stream.State == types.StreamStateStreamLive {
			metadata := strconv.FormatInt(stream.ViewerCount, 10)
			if err := client.PutMetadata(ctx, channelArn, metadata); err != nil {
				return nil
			}
		}
		time.Sleep(time.Second)
	}

	return nil
}

func createChannel(ctx context.Context, client *ivsclient.Client) error {
	channel, stream, err := client.CreateChannel(ctx, "test")
	if err != nil {
		return err
	}
	fmt.Printf("Channel arn: %v\nIngest endpoint: %v\nPlayback url: %v\n", *channel.Arn, *channel.IngestEndpoint, *channel.PlaybackUrl)
	fmt.Printf("Stream key: %v\n", *stream.Value)
	return nil
}

func deleteChannel(ctx context.Context, client *ivsclient.Client) error {
	if channelArn == nil || len(*channelArn) == 0 {
		return errors.New("Empty channel arn")
	}
	return client.DeleteChannel(ctx, channelArn)
}

func run(ctx context.Context) error {
	initFlags()
	if operation == nil {
		return errors.New("Operation was not provided")
	}

	client, err := ivsclient.NewClient(ctx)
	if err != nil {
		return err
	}

	switch *operation {
	case "createChannel":
		return createChannel(ctx, client)
	case "deleteChannel":
		return deleteChannel(ctx, client)
	case "streamOnlineViewers":
		return streamOnlineViewers(ctx, client)
	default:
		return errors.New(fmt.Sprintf("Unsupported operation: %q", *operation))
	}
}
func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
