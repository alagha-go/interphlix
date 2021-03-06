package casts

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//27cc73002943ca37c5422425af69c720//


type Cast struct {
	ID									*primitive.ObjectID				`json:"_id,omitempty" bson:"_id,omitempty"`
	Name               					string      					`json:"name,omitempty" bson:"name,omitempty"`
	AlsoKnownAs        					[]string    					`json:"also_known_as,omitempty" bson:"also_known_as,omitempty"`
	Biography          					string      					`json:"biography,omitempty" bson:"biography,omitempty"`
	Birthday           					string      					`json:"birthday,omitempty" bson:"birthday,omitempty"`
	Deathday           					string 							`json:"deathday,omitempty" bson:"deathday,omitempty"`
	Gender             					int64       					`json:"gender,omitempty" bson:"gender,omitempty"`
	Homepage           					interface{} 					`json:"homepage,omitempty" bson:"homepage,omitempty"`
	TmdbID                 				int64       					`json:"tmdb_id,omitempty" bson:"tmdb_id,omitempty"`
	ImdbID             					string      					`json:"imdb_id,omitempty" bson:"imdb_id,omitempty"`
	KnownForDepartment 					string      					`json:"known_for_department,omitempty" bson:"known_for_department,omitempty"`
	PlaceOfBirth       					string      					`json:"place_of_birth,omitempty" bson:"place_of_birth,omitempty"` 
	Popularity         					float64     					`json:"popularity,omitempty" bson:"popularity,omitempty"`
	ProfilePath        					string      					`json:"profile_path,omitempty" bson:"profile_path,omitempty"`
}


type Response struct {
	Page         						int64    						`json:"page"`
	Results      						[]Result 						`json:"results"`
	TotalPages   						int64    						`json:"total_pages"`
	TotalResults 						int64    						`json:"total_results"`
}

type Result struct {
	Adult              					bool          					`json:"adult"`
	Gender             					int64         					`json:"gender"`
	ID                 					int64         					`json:"id"`
	KnownFor           					[]interface{} 					`json:"known_for"`
	KnownForDepartment 					string        					`json:"known_for_department"`
	Name               					string        					`json:"name"`
	Popularity         					float64       					`json:"popularity"`
	ProfilePath        					string        					`json:"profile_path"`
}
