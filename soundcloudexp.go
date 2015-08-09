package soundcloudexp

import (
	"fmt"
	"encoding/json"
	"github.com/heatxsink/go-httprequest"
)

const (
	SC_REPOST_URL = "https://api-v2.soundcloud.com/profile/soundcloud:users:%d?limit=%d&offset=%d"
	SC_OEMBED_URL = "http://soundcloud.com/oembed"
)

type RepostReponse struct {
	Collection []Item `json:"collection,omitempty"`
	NextHref string `json:"next_href,omitempty"`
}

type Item struct {
	Type string `json:"type,omitempty"`
	User User `json:"user,omitempty"`
	Track Track `json:"track,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	CreateAt string `json:"created_at,omitempty"`
}

type User struct {
	FullName string `json:"full_name,omitempty"`
	Country string `json:"country,omitempty"`
	City string `json:"city,omitempty"`
	TrackCount uint64 `json:"track_count,omitempty"`
	FollowersCount uint64 `json:"followers_count,omitempty"`
	FollowingsCount uint64 `json:"followings_count,omitempty"`
	PublicFavoritesCount uint64 `json:"public_favorites_count,omitempty"`
	GroupsCount uint64 `json:"group_count,omitempty"`
	Description string `json:"description,omitempty"`
	Plan string `json:"plan,omitempty"`
	Id uint64 `json:"id,omitempty"`
	Uri string `json:"uri,omitempty"`
	Username string `json:"username,omitempty"`
	Kind string `json:"kind,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	PermalinkUrl string `json:"permalink_url,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
}

type Track struct {
	User User `json:"user,omitempty"`
	UserId uint64 `json:"user_id,omitempty"`
	Genre string  `json:"genre,omitempty"`
	TagList string  `json:"tag_list,omitempty"`
	Duration uint64  `json:"duration,omitempty"`
	Downloadable bool  `json:"downloadable,omitempty"`
	Streamable bool  `json:"streamable,omitempty"`
	OriginalContentSize uint64  `json:"original_content_size,omitempty"`
	Commentable bool `json:"commentable,omitempty"`
	Sharing string `json:"sharing,omitempty"`
	Public bool `json:"public,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	PermalinkUrl string `json:"permalink_url,omitempty"`
	Kind string `json:"kind,omitempty"`
	Id uint64 `json:"id,omitempty"`
	ArtworkUrl string `json:"artwork_url,omitempty"`
	Uri string `json:"uri,omitempty"`
	StreamUrl string `json:"stream_url,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type OEmbedResponse struct {
	Version float64 `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	ProviderUrl string `json:"provider_url,omitempty"`
	Height uint64 `json:"height,omitempty"`
	Width string `json:"width,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
	Html string `json:"html,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	AuthorUrl string `json:"author_url,omitempty"`
}

func UserReposts(user_id uint64, limit uint64, offset uint64) (RepostReponse, error) {
	var response RepostReponse
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
	url := fmt.Sprintf(SC_REPOST_URL, user_id, limit, offset)
	hr := httprequest.NewWithDefaults()
	body, status_code, err := hr.Get(url, headers)
	if err != nil {
		return response, err
	}
	if status_code == 200 {
		err := json.Unmarshal(body, &response)
		if err != nil {
			return response, err
		}
	}
	return response, err
}

func OEmbed(url string) (OEmbedResponse, error) {
	return OEmbedWithOptions(url, "CCC", "json")
}

func OEmbedWithOptions(stream_url string, color string, format string) (OEmbedResponse, error) {
	var response OEmbedResponse
	url := fmt.Sprintf("%s?iframe=true&show_comments=false&color=%s&format=%s&url=%s", SC_OEMBED_URL, color, format, stream_url)
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
	hr := httprequest.NewWithDefaults()
	body, status_code, err := hr.Get(url, headers)
	if err != nil {
		return response, err
	}
	if status_code == 200 {
		err := json.Unmarshal(body, &response)
		if err != nil {
			return response, err
		}
	}
	return response, err
}
