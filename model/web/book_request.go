package web

type BookCreateRequest struct {
	Title        string `json:"title" validate:"required"`
	Publisher    string `json:"publisher" validate:"required"`
	BookType     string `json:"book_type" validate:"required"`
	YearReleased string `json:"year_released" validate:"required"`
	Synopsis     string `json:"synopsis" validate:"required"`
	Genre        string `json:"genre" validate:"required"`
	Stock        uint64 `json:"stock" validate:"required"`
	WriterId     uint64 `json:"writer_id,omitempty"`
}

type BookUpdateRequest struct {
	Id           string `json:"id" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Publisher    string `json:"publisher" validate:"required"`
	BookType     string `json:"book_type" validate:"required"`
	YearReleased string `json:"year_released" validate:"required"`
	Synopsis     string `json:"synopsis" validate:"required"`
	Genre        string `json:"genre" validate:"required"`
	Stock        uint64 `json:"stock" validate:"required"`
	WriterId     uint64 `json:"writer_id,omitempty"`
}
