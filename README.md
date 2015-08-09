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
	response, err := soundcloudexp.UserReposts(20804924, 30, 0)
	for _, item := range response.Collection {
		if item.Type == "track-repost" {
			oembed, err := soundcloudexp.OEmbed(item.Track.Uri)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(oembed.Html)
		}
	}
}

```
