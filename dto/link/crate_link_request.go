package link

type CreateLinkRequest struct {
	Name string `json:"name" validate:"required,min=2,max=255,alphanum"`
	Url  string `json:"url"  validate:"required,url"`
}
