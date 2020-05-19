package models

type Seed struct {
	Url       string `json:"url" form:"url" query:"url"`
	Title     string `json:"title" form:"title" query:"title"`
	Content   string `json:"content" form:"content" query:"content"`
	OneParent string `json:"oneParent" form:"oneParent" query:"oneParent"`
}
