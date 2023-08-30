package domain

type PostSlice struct {
	Text string
	Url  string
}

func NewPostSlice(text string, url string) PostSlice {
	return PostSlice{
		Text: text,
		Url:  url,
	}
}
