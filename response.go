package ga

import (
	"encoding/csv"
	"io"
)

//HelperResponse struct
type HelperResponse struct {
	Reports []Response `json:"reports"`
}

func (h *HelperResponse) headers() (headers []string) {
	for _, dim := range h.Reports[0].ColumnHeader.Dimensions {
		headers = append(headers, dim)
	}

	for _, metric := range h.Reports[0].ColumnHeader.MetricHeader.MetricHeaderEntries {
		headers = append(headers, metric.Name)
	}

	return
}

func (h *HelperResponse) process() (res []map[string]interface{}) {

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
func (h *HelperResponse) Export() []map[string]interface{} {
	return h.process()
}

//CSV export
func (h *HelperResponse) CSV(out io.Writer) error {
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
