package web

type BookCreateRequest struct {
	Title        string `json:"title" binding:"required"`
	Publisher    string `json:"publisher" binding:"required"`
	BookType     string `json:"book_type" binding:"required"`
	YearReleased string `json:"year_released" binding:"required"`
	Synopsis     string `json:"synopsis" binding:"required"`
	Genre        string `json:"genre" binding:"required"`
	Stock        uint64 `json:"stock" binding:"required"`
	WriterId     uint64 `json:"writer_id,omitempty" binding:"required"`
}

type BookUpdateRequest struct {
	Id           uint64
	Title        string `json:"title" binding:"required"`
	Publisher    string `json:"publisher" binding:"required"`
	BookType     string `json:"book_type" binding:"required"`
	YearReleased string `json:"year_released" binding:"required"`
	Synopsis     string `json:"synopsis" binding:"required"`
	Genre        string `json:"genre" binding:"required"`
	Stock        uint64 `json:"stock" binding:"required"`
	WriterId     uint64 `json:"writer_id,omitempty" binding:"required"`
}
