package dao

const (
	asc  = "asc"
	desc = "desc"
)

func sortOrder(ok bool) string {
	if !ok {
		return desc
	}

	return asc
}
