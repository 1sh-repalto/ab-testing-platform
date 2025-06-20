package utils

import (
	"backend/models"
	"math/rand"
)

func PickWeightedVariant(variants []models.Variant) int {
	var totalWeight float64
	for _, v := range variants {
		totalWeight += v.Weight
	}

	r := rand.Float64() * totalWeight
	var cumulative float64

	for _, v := range variants {
		cumulative += v.Weight
		if r < cumulative {
			return v.ID
		}
	}

	// fallback
	return variants[0].ID
}
