package domain

type OpenIAEmbedder interface {
	Embed(slices []PostSlice) ([]Vector, error)
}
