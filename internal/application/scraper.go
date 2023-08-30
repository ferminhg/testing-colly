package application

import (
	"github.com/ferminhg/testing-colly/internal/domain"
)

type Scrapper struct {
	strategies []domain.SiteStrategy
}

func NewScrapper(strategies []domain.SiteStrategy) *Scrapper {
	return &Scrapper{
		strategies: strategies,
	}
}

func (s *Scrapper) Run() error {
	for _, strategy := range s.strategies {
		if err := strategy.Execute(); err != nil {
			return err
		}
		if err := strategy.MarshalPosts(); err != nil {
			return err
		}
	}
	return nil
}
