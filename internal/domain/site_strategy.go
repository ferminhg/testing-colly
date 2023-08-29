package domain

type SiteStrategy interface {
	Execute() error
	Marshal(file string) error
}
