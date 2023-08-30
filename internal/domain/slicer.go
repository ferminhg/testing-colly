package domain

import (
	"log"
	"regexp"
)

type Slicer struct {
}

func NewSlicer() *Slicer {
	return &Slicer{}
}

func (ps *Slicer) Slice(p Post) []PostSlice {
	log.Println("✂️ Slicing post", p.Id)
	var slices []PostSlice
	slices = append(slices, NewPostSlice(p.Title, p.Link.String()))
	a := regexp.MustCompile(`\n`)
	for _, paragraph := range a.Split(p.Text, -1) {
		slices = append(slices, NewPostSlice(paragraph, p.Link.String()))
	}
	log.Println("✂️ Num Slices", len(slices))
	return slices
}
