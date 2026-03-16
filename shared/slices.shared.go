package shared

type sliceUtils struct{}

func NewSliceUtils() *sliceUtils {
	return &sliceUtils{}
}

func (h *sliceUtils) DifferenceString(a, b []string) []string {
	setB := make(map[string]bool)
	for _, v := range b {
		setB[v] = true
	}

	var diff []string
	for _, v := range a {
		if !setB[v] {
			diff = append(diff, v)
		}
	}

	if len(diff) == 0 {
		return make([]string, 0)
	}

	return diff
}

func (h *sliceUtils) IntersectionString(a, b []string) []string {
	setB := make(map[string]bool)
	for _, v := range b {
		setB[v] = true
	}

	var intersection []string
	for _, v := range a {
		if setB[v] {
			intersection = append(intersection, v)
		}
	}

	if len(intersection) == 0 {
		return make([]string, 0)
	}

	return intersection
}

func (h *sliceUtils) ContainsString(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
