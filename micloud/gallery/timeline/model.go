package timeline

type Timeline struct {
	// eg: 20240101 : 1
	DayCount  map[int]int `json:"dayCount"`
	IndexHash int64       `json:"indexHash"`
}
