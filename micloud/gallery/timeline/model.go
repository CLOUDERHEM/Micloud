package timeline

type Timeline struct {
	// eg: 20240101 : 1
	DayCount  map[string]int `json:"dayCount"`
	IndexHash int64          `json:"indexHash"`
}
