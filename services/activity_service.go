package services

import (
	"context"
	"encoding/json"
	"strconv"

	"bitbucket.com/devplaylister/playlister-activity/config"
	"bitbucket.com/devplaylister/playlister-activity/dto"
	"github.com/olivere/elastic"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type UserSongActivityIndex struct {
	UserID string   `json:"user_id"`
	Artist string   `json:"artist"`
	Album  string   `json:"album"`
	Track  string   `json:"track"`
	Loc    Location `json:"location"`
}

func FindNearest(lat float64, lon float64, distance int) (interface{}, error) {
	client, _ := config.GetClient()
	distanceQuery := elastic.NewGeoDistanceQuery("location")
	distanceQuery = distanceQuery.Lat(lat)
	distanceQuery = distanceQuery.Lon(lon)
	distanceQuery = distanceQuery.Distance(strconv.Itoa(distance) + "m")
	boolQuery := elastic.NewBoolQuery()
	boolQuery = boolQuery.Must(elastic.NewMatchAllQuery())
	boolQuery = boolQuery.Filter(distanceQuery)
	result, err := client.Search().Index("user_current_activity").Type("song_activity").Query(boolQuery).Do(context.Background())

	if err != nil {
		return nil, err
	}
	return extractActivities(result)
}

func extractActivities(result *elastic.SearchResult) ([]UserSongActivityIndex, error) {
	activities := make([]UserSongActivityIndex, 0)
	var err error
	if result.Hits.TotalHits > 0 {
		for _, hit := range result.Hits.Hits {
			var activity UserSongActivityIndex
			err = json.Unmarshal(*hit.Source, &activity)
			if err != nil {
				break
			}

			activities = append(activities, activity)
		}
	}

	return activities, err
}

func IndexActivity(activity dto.UserSongActivityData) error {
	client, _ := config.GetClient()

	activityIndex := UserSongActivityIndex{
		UserID: strconv.Itoa(activity.UserID),
		Artist: activity.Artist,
		Album:  activity.Album,
		Track:  activity.Track,
		Loc: Location{
			Lat: activity.Latitude,
			Lon: activity.Longitude,
		},
	}
	client.CreateIndex("user_current_activity").BodyString(mapping).Do(context.Background())
	indexActivityRequest := elastic.NewBulkIndexRequest().Index("user_current_activity").Type("song_activity").Id(activityIndex.UserID).Doc(activityIndex)
	bulkRequest := client.Bulk()
	bulkRequest = bulkRequest.Add(indexActivityRequest)

	bulkResponse, _ := bulkRequest.Do(context.Background())
	err := config.CheckBulkResponse(bulkResponse)

	return err
}

const mapping = `{
	"mappings":{
		"song_activity":{
			"properties":{
				"user_id":{
					"type":"keyword"
				},
				"artist":{
					"type":"keyword"
				},
				"album":{
					"type":"keyword"
				},
				"track":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				}
			}
		}
	}
}`
