package movies

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Movie struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Code										string									`json:"code,omitempty" bson:"code,omitempty"`
	Title										string									`json:"title,omitempty" bson:"title,omitempty"`
	Type										string									`json:"type,omitempty" bson:"type,omitempty"`
    Trending                                    bool                                    `json:"trending,omitempty" bson:"trending,omitempty"`
    Featured                                    bool                                    `json:"featured,omitempty" bson:"featured,omitempty"`
    Popular                                     bool                                    `json:"popular,omitempty" bson:"popular,omitempty"`
    Available                                   bool                                    `json:"available,omitempty" bson:"available,omitempty"`
    PageUrl										string									`json:"page_url,omitempty" bson:"page_url,omitempty"`
    ImageUrl									string									`json:"image_url,omitempty" bson:"image_url,omitempty"`
    Released									*time.Time								`json:"released,omitempty" bson:"released,omitempty"`
    Urls										[]string								`json:"urls,omitempty" bson:"urls,omitempty"`
    Genres										[]string								`json:"genre,omitempty" bson:"genre,omitempty"`
    Views                                       int                                     `json:"views,omitempty" bson:"views,omitempty"`
    Rating                                      float64                                 `json:"rating,omitempty" bson:"rating,omitempty"`
    Raters                                      int                                     `json:"raters,omitempty" bson:"raters,omitempty"`
	Server										*Server									`json:"server,omitempty" bson:"server,omitempty"`
	DateUpdated									*time.Time								`json:"date_updated,omitempty" bson:"date_updated,omitempty"`
    Servers										[]Server								`json:"servers,omitempty" bson:"servers,omitempty"`
    Seasons										[]Season								`json:"seasons,omitempty" bson:"seasons,omitempty"`
    Casts										[]string								`json:"casts,omitempty" bson:"casts,omitempty"`
    Duration									time.Duration							`json:"duration,omitempty" bson:"duration,omitempty"`
    Countries									[]string								`json:"countries,omitempty" bson:"countries,omitempty"`
    Producers									[]string								`json:"producers,omitempty" bson:"producers,omitempty"`
    Percentage					                float64									`json:"percentage,omitempty" bson:"percentage,omitempty"`
    Description									string									`json:"description,omitempty" bson:"description,omitempty"`
}


type Season struct {
    ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Name										string									`json:"name,omitempty" bson:"name,omitempty"`
    Index                                       int                                     `json:"index,omitempty" bson:"index,omitempty"`
    Code										string									`json:"code,omitempty" bson:"code,omitempty"`
	Episodes									[]Episode								`json:"episodes,omitempty" bson:"episodes,omitempty"`
}


type Episode struct {
    ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Name										string									`json:"name,omitempty" bson:"name,omitempty"`
    Index                                       int                                     `json:"index,omitempty" bson:"index,omitempty"`
    Available                                   bool                                    `json:"available,omitempty" bson:"available,omitempty"`
    Code										string									`json:"code,omitempty" bson:"code,omitempty"`
    Views                                       int                                     `json:"views,omitempty" bson:"views,omitempty"`
    Urls										[]string								`json:"urls,omitempty" bson:"urls,omitempty"`
	Server										*Server									`json:"server,omitempty" bson:"server,omitempty"`
	Servers										[]Server								`json:"servers,omitempty" bson:"servers,omitempty"`
    Percentage					                float64									`json:"percentage,omitempty" bson:"percentage,omitempty"`
}


type Server struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
    Name                                        string                               	`json:"name,omitempty" bson:"name,omitempty"`
    Id                                          string                               	`json:"id,omitempty" bson:"id,omitempty"`
    WatchID                                     string                               	`json:"watch_id,omitempty" bson:"watch_id,omitempty"`
    Url                                         string                               	`json:"url,omitempty" bson:"url,omitempty"`
}


type Recommendation struct {
    Seed                                        int64                                   `json:"seed,omitempty" bson:"seed,omitempty"`
    Categories                                  []Category                              `json:"categories,omitempty" bson:"categories,omitempty"`
}


type Category struct {
    Title                                       string                                  `json:"title,omitempty" bson:"title,omitempty"`
    Path                                        string                                  `json:"path,omitempty" bson:"path,omitempty"`
    Movies                                      []Movie                                 `json:"movies,omitempty" bson:"movies,omitempty"`
}