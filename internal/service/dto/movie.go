package dto

type MovieCreateRequest struct {
	Title string `json:"title"`
}

type MovieSingleResponse struct {
	ID              int64   `json:"id"`
	Title           string  `json:"title"`
	EpisodeDuration *int32  `json:"episode_duration"`
	EpisodeCount    *int16  `json:"episode_count"`
	AlterTitles     *string `json:"alter_titles"`
	Description     *string `json:"description"`
}

type MovieItemData struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}

type MovieListResponse struct {
	Movies []MovieItemData `json:"movies"`
}

type MovieUpdateRequest struct {
	Title           *string `json:"title"`
	EpisodeDuration *int32  `json:"episode_duration"`
	EpisodeCount    *int16  `json:"episode_count"`
	AlterTitles     *string `json:"alter_titles"`
	Description     *string `json:"description"`
}
