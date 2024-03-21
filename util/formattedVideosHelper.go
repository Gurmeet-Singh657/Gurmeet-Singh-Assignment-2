package util

import (
	"Assignment/model"
	"time"
	"google.golang.org/api/youtube/v3"
)

// formats the ideo to storable content
func GetFormattedVideos(response *youtube.SearchListResponse) ([]model.Video, error){

	var videos []model.Video
	for _, item := range response.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			return nil,err
		}

		video := model.Video{
			ID:          item.Id.VideoId,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishedAt: publishedAt.Format(time.RFC3339),
			Thumbnail: model.Thumbnail{
				Default: item.Snippet.Thumbnails.Default.Url,
				Medium:  item.Snippet.Thumbnails.Medium.Url,
				High:    item.Snippet.Thumbnails.High.Url,
			},
		}
		videos = append(videos, video)
	}

	return videos,nil
}