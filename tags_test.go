package gomicSocMedShared

import "testing"

func TestGetArtTagsInRandomOrder(t *testing.T) {
	tags := GetArtTagsInRandomOrder()
	counter := 0
	for _, t := range TAGS {
		if contains(tags, t) {
			counter += 1
		}
	}
	if len(TAGS) != counter {
		t.Errorf("Expected slice of randomized tags to contain %v elements, but it contained %v.", len(TAGS), counter)
	}
}

func TestEnrichArtTags(t *testing.T) {
	tags := GetArtTagsInRandomOrder()
	specificTags := []string{"specific1", "specific2", "specific3"}

	enrichedTags := EnrichArtTags(specificTags)

	expected := len(tags) + 3
	actual := len(enrichedTags)

	if actual != expected {
		t.Error("Expected tag slice length to be", expected, "but it was", actual)
	}
}

func contains(collection []string, item string) bool {
	for _, ci := range collection {
		if ci == item {
			return true
		}
	}
	return false
}
