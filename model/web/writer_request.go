package web

type WriterCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type WriterUpdateRequest struct {
	Id   uint64 `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
