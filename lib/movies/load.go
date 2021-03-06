package movies

import (
	"context"
	"errors"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// load featured movies
func LoadFeaturedMovies(start, end int, seed int64, Type ...string) []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.M{"_id": 1, "image_url": 1, "title": 1, "type": 1,})

	filter := bson.M{"featured": true}

	if len(Type) > 0 {
		filter = bson.M{"type": Type[0], "featured": true}
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadFeaturedMovies")
		return Movies
	}
	cursor.All(ctx, &Movies)

	if seed != 0 {
		Movies = RandomMovies(seed, Movies)
	}

	if start > len(Movies) {
		return []Movie{}
	}else if end > len(Movies) {
		return Movies[start:]
	}

	return Movies[start:end]
}

// loading trending Movies
func LoadTrendingMovies(start, end int, seed int64, Type ...string) []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.M{"_id": 1, "image_url": 1, "title": 1, "type": 1,})

	filter := bson.M{"trending": true}

	if len(Type) > 0 {
		filter = bson.M{"type": Type[0], "trending": true}
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadTrendingMovies")
		return Movies
	}
	cursor.All(ctx, &Movies)

	if seed != 0 {
		Movies = RandomMovies(seed, Movies)
	}

	if start > len(Movies) {
		return []Movie{}
	}else if end > len(Movies) {
		return Movies[start:]
	}

	return Movies[start:end]
}

// load movies by genre
func LoadMoviesByGenre(genre string) []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.M{"_id": 1, "image_url": 1, "title": 1, "type": 1,})

	cursor, err := collection.Find(ctx, bson.M{"genre": genre}, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadMoviesByGenre")
		return Movies
	}
	cursor.All(ctx, &Movies)

	return Movies
}

// load movies by type and genre
func LoadMoviesByGenreAndType(Type, genre string) []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.M{"_id": 1, "image_url": 1, "title": 1, "type": 1,})

	cursor, err := collection.Find(ctx, bson.M{"type": Type , "genre": genre}, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadMoviesByGenre")
		return Movies
	}
	cursor.All(ctx, &Movies)

	return Movies
}

// load popular content
func LoadPopulareContent(start, end int, seed int64, Type ...string) []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	opts := options.Find().SetProjection(bson.M{"_id": 1, "image_url": 1, "title": 1, "type": 1,})

	filter := bson.M{"popular": true}

	if len(Type) > 0 {
		filter = bson.M{"type": Type[0], "popular": true}
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		variables.SaveError(err, "movies", "LoadPopularTvShows")
		return Movies
	}
	cursor.All(ctx, &Movies)

	if seed != 0 {
		Movies = RandomMovies(seed, Movies)
	}

	if start > len(Movies) {
		return []Movie{}
	}else if end > len(Movies) {
		return Movies[start:]
	}

	return Movies[start:end]
}

// load movie contents
func LoadMovie(ID primitive.ObjectID) (*Movie, error) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	var Movie Movie

	opts := options.FindOne().SetProjection(bson.M{"page_url": 0, "server": 0, "servers": 0, "seasons": 0, "code": 0})

	err := collection.FindOne(ctx, bson.M{"_id": ID}, opts).Decode(&Movie)
	if err != nil {
		return nil, errors.New(variables.MovieNotFound)
	}
	return &Movie, nil
}