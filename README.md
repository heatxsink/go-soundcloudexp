go-soundcloudexp
===
An experimental soundcloud api module (for an undocumented api of course!).

Install
-------
You'll need to have golang installed.

	$ go get github.com/heatxsink/go-soundcloudexp

Usage
-----

```go

package main

import (
  "fmt"
  "github.com/heatxsink/go-soundcloudexp"
)

func main() {
	user_id := uint64(20804924)
	limit := uint64(10)
	offset := uint64(0)
	response, _ := soundcloudexp.UserReposts(user_id, limit, offset)
	for _, item := range response.Collection {
		if item.Type == "track-repost" {
			oembed, _ := soundcloudexp.OEmbed(item.Track.Uri)
			fmt.Println(oembed.Html)
		}
	}
}

```
