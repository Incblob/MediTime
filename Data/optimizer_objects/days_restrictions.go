package optimizerobjects

type days_object_restrictions interface {
	badness_score(d Day) float32
	validate() (bool, string)
}
