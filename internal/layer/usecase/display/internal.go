package display

func ellipsis(txt string, limit int) string {
	if len(txt) > limit-3 {
		return txt[:limit-3] + "..."
	}

	return txt
}
