package usemergeconcurrent

type Pair struct {
	Key   int64
	Value float64
}

type InitInfo struct {
	InputFile  string
	OutputFile string
	CPU        int64
}
