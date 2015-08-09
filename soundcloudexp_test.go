package soundcloudexp

import (
	"fmt"
	"testing"
)

func TestUserReposts(t *testing.T) {
	response, _ := UserReposts(20804924, 30, 0)
	for _, item := range response.Collection {
		if item.Type == "track-repost" {
			oembed, err := OEmbed(item.Track.Uri)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(oembed.Html)
		}
	}
}