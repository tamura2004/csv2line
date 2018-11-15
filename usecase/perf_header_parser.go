package usecase

import (
	"github.com/tamura2004/csv2line/domain"
)

type PerfHeaderParser struct {
	header []domain.Header
}

func (p *PerfHeaderParser) ParseHeader(row []string) {
	p.header = make([]domain.Header, len(row)-1)

	for j, header := range row[1:] {
		p.header[j] = domain.NewHeader(header)
	}
}

func (p *PerfHeaderParser) Header(i int) domain.Header {
	return p.header[i]
}
