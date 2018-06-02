package gomicSocMedShared

import "math/rand"

const (
	MAX_LENGTH = 140
)

var TAGS = []string{"art", "artist", "artwork", "bandedesinée", "cartoon", "comic", "comicart", "comics", "comix", "concept", "conceptart", "create", "creative", "dessin", "design", "draw", "drawing", "drawings", "fiction", "graphicnovel", "illustration", "image", "imagination", "inked", "inking", "manga", "painting", "sketch", "sketchbook", "sequentialart", "zeichnung", "挿絵", "絵", "イラスト", "일러스트", "일러스트레이션", "アート", "マンガ", "イラストレーション", "ドローイング", "スケッチ", "ilustração", "ilustració", "peinture"}

func GetArtTagsInRandomOrder() []string {
	shuffeled := make([]string, len(TAGS))
	perm := rand.Perm(len(TAGS))
	for i, v := range perm {
		shuffeled[v] = TAGS[i]
	}
	return shuffeled
}

func EnrichArtTags(tags []string) []string {
	for _, tag := range GetArtTagsInRandomOrder() {
		tags = append(tags, tag)
	}
	return tags
}
