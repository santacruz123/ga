package ga

import (
	"encoding/csv"
	"io"
)

//Response struct
type Response struct {
	Reports []response `json:"reports"`
}

func (h *Response) headers() (headers []string) {
	for _, dim := range h.Reports[0].ColumnHeader.Dimensions {
		headers = append(headers, dim)
	}

	for _, metric := range h.Reports[0].ColumnHeader.MetricHeader.MetricHeaderEntries {
		headers = append(headers, metric.Name)
	}

	return
}

func (h *Response) process() (res []map[string]interface{}) {

	for _, dataRow := range h.Reports[0].Data.Rows {
		row := make(map[string]interface{})

		for i, dim := range h.Reports[0].ColumnHeader.Dimensions {
			row[dim] = dataRow.Dimensions[i]
		}

		for i, metric := range h.Reports[0].ColumnHeader.MetricHeader.MetricHeaderEntries {
			row[metric.Name] = dataRow.Metrics[0].Values[i]
		}

		res = append(res, row)
	}

	return
}

//Export structure
func (h *Response) Export() []map[string]interface{} {
	return h.process()
}

//CSV export
func (h *Response) CSV(out io.Writer) error {
	data := h.process()

	w := csv.NewWriter(out)

	headers := h.headers()

	if err := w.Write(headers); err != nil {
		return err
	}

	for _, record := range data {
		row := []string{}
		for i := range headers {
			row = append(row, record[headers[i]].(string))
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}

	w.Flush()
	return w.Error()
}
