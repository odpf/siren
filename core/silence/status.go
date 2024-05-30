package silence

const (
	StatusTotal   = "total"
	StatusPartial = "partial"
)

func Status(hasSilenced, hasNonSilenced bool) string {
	if hasSilenced && !hasNonSilenced {
		return StatusTotal
	} else if hasSilenced && hasNonSilenced {
		return StatusPartial
	}
	return ""
}
