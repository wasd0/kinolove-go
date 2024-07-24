package mapper

import (
	"fmt"
	. "kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/service/dto"
)

func MapMovieToSingleResponse(movie *Movies) dto.MovieSingleResponse {
	if movie == nil {
		return dto.MovieSingleResponse{}
	}

	return dto.MovieSingleResponse{
		ID:              movie.ID,
		Title:           movie.Title,
		EpisodeDuration: movie.EpisodeDuration,
		EpisodeCount:    movie.EpisodeCount,
		AlterTitles:     movie.AlterTitles,
		Description:     movie.Description,
	}
}

func MapMovieToItemData(movie *Movies) dto.MovieItemData {
	if movie == nil {
		return dto.MovieItemData{}
	}

	return dto.MovieItemData{Title: movie.Title}
}

func MapUpdateRequestToMovie(request *dto.MovieUpdateRequest, movie *Movies) error {
	if request == nil {
		return fmt.Errorf("movie update request is nil")
	}

	if request.Title != nil {
		movie.Title = *request.Title
	}

	if request.EpisodeDuration != nil {
		movie.EpisodeDuration = request.EpisodeDuration
	}

	if request.AlterTitles != nil {
		movie.AlterTitles = request.AlterTitles
	}

	if request.Description != nil {
		movie.Description = request.Description
	}

	return nil
}
