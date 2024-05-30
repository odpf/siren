package structure

import (
	"crypto/sha256"
	"fmt"

	"github.com/mitchellh/hashstructure/v2"
)

func RemoveDuplicate[T comparable](aSlice []T) []T {
	keys := make(map[T]bool)
	list := []T{}

	for _, v := range aSlice {
		if _, value := keys[v]; !value {
			keys[v] = true
			list = append(list, v)
		}
	}
	return list
}

func GroupByLabels[S any](collection []S, groupBy []string, getLabels func(S) map[string]string) (map[uint64][]S, error) {
	var collectionMap = map[uint64][]S{}

	for _, c := range collection {
		var labels = getLabels(c)
		var groupLabels = BuildGroupLabels(labels, groupBy)
		if len(groupLabels) == 0 {
			groupLabels = labels
		}
		hash, err := hashstructure.Hash(groupLabels, hashstructure.FormatV2, nil)
		if err != nil {
			return nil, fmt.Errorf("cannot get hash from alert %v", c)
		}
		collectionMap[hash] = append(collectionMap[hash], c)
	}

	return collectionMap, nil
}

func BuildGroupLabels(labels map[string]string, groupBy []string) map[string]string {
	var groupLabels = map[string]string{}

	for _, g := range groupBy {
		if v, ok := labels[g]; ok {
			groupLabels[g] = v
		}
	}

	return groupLabels
}

// HashGroupKey hash groupKey from alert and hashKey from labels
func HashGroupKey(groupKey string, hashKey uint64) string {
	h := sha256.New()
	// hash.Hash.Write never returns an error.
	//nolint: errcheck
	h.Write([]byte(fmt.Sprintf("%s%d", groupKey, hashKey)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
