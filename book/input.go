package book

type BookInput struct {
	Price int    `json:"price" binding:"required,number"`
	Title string `json:"title" binding:"required"`
}
