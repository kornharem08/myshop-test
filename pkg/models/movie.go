package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Viewer struct {
	Rating     float64 `bson:"rating,omitempty" json:"rating,omitempty"`
	NumReviews int     `bson:"numReviews,omitempty" json:"numReviews,omitempty"`
	Meter      int     `bson:"meter,omitempty" json:"meter,omitempty"`
}

type Tomatoes struct {
	Dvd         time.Time `bson:"dvd,omitempty" json:"dvd,omitempty"`
	LastUpdated time.Time `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Viewer      Viewer    `bson:"viewer,omitempty" json:"viewer,omitempty"`
}

type Movie struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Plot     string             `json:"plot,omitempty" bson:"plot,omitempty"`
	Genres   []string           `json:"genres,omitempty" bson:"genres,omitempty"`
	Runtime  int                `json:"runtime,omitempty" bson:"runtime,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Tomatoes Tomatoes           `json:"tomatoes,omitempty" bson:"tomatoes,omitempty"`
	Year     int                `json:"year,omitempty" bson:"year,omitempty"`
	Poster   string             `json:"poster,omitempty" bson:"poster,omitempty"`
}
