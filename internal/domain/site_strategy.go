package domain

type SiteStrategy interface {
	Execute() error
	MarshalPosts() error
}
