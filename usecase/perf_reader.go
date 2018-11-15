package usecase

type Scanner interface {
	Scan() bool
	Record() []string
	LineNumber() int
}

type PerfReader struct {
	Scanner
}

func NewPerfReader(s Scanner) *PerfReader {
	return &PerfReader{
		Scanner: s,
	}
}

func (r *PerfReader) Copy(w *PerfWriter) {
	for r.Scan() {
		w.Write(r.Record(), r.LineNumber())
	}
}
