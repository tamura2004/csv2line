package main

func parseHeaderRow(row []string) []TagFieldKey {
	tagFieldKeys := make([]TagFieldKey, len(row)-1)

	for j, header := range row[1:] {
		keys, err := parseHeader(header)
		handleError(err)
		tagFieldKeys[j] = keys
	}
	return tagFieldKeys
}
