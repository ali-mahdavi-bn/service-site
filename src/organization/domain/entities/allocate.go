package entities

import (
	"sort"

	"github.com/pkg/errors"
)

func Allocate(line OrderLine, batches ...*Batch) (reference string, err error) {
	sort.Slice(batches, func(i, j int) bool {
		leftTime := batches[i].ETA
		rightTime := batches[j].ETA

		return leftTime.Time.Before(rightTime.Time)
	})
	for _, batch := range batches {
		if batch.CanAllocate(line) {
			batch.Allocate(line)
			return batch.Reference, nil
		}
	}
	return "", errors.Errorf("Out of stock for SKU %s", line.SKU)
}
