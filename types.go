package ga

import "os"

var access = os.Getenv("ACCESS")

//MetricType type
type MetricType string

//MetricType conststants
const (
	MetricTypeUnspecified MetricType = "METRIC_TYPE_UNSPECIFIED"
	MetricTypeInteger     MetricType = "INTEGER"
	MetricTypeFloat       MetricType = "FLOAT"
	MetricTypeCurrency    MetricType = "CURRENCY"
	MetricTypePercent     MetricType = "PERCENT"
	MetricTypeTime        MetricType = "TIME"
)

//DateRange struct
type DateRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

//Dimension struct
type Dimension struct {
	Name             string   `json:"name"`
	HistogramBuckets []string `json:"histogramBuckets,omitempty"`
}

//Metric struct
type Metric struct {
	Expression     string     `json:"expression"`
	Alias          string     `json:"alias,omitempty"`
	FormattingType MetricType `json:"formattingType,omitempty"`
}

//Request struc
type Request struct {
	ViewID     string      `json:"viewId"`
	DateRange  []DateRange `json:"dateRanges"`
	Dimensions []Dimension `json:"dimensions,omitempty"`
	Metrics    []Metric    `json:"metrics,omitempty"`
	PageToken  string      `json:"pageToken,omitempty"`
	PageSize   int64       `json:"pageSize,omitempty"`
}

//Response type
type Response struct {
	ColumnHeader struct {
		Dimensions   []string `json:"dimensions"`
		MetricHeader struct {
			MetricHeaderEntries []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"metricHeaderEntries"`
		} `json:"metricHeader"`
	} `json:"columnHeader"`

	Data struct {
		Rows []struct {
			Dimensions []string `json:"dimensions"`
			Metrics    []struct {
				Values []string `json:"values"`
			} `json:"metrics"`
		} `json:"rows"`

		RowCount     int64 `json:"rowCount"`
		IsDataGolden bool  `json:"isDataGolden"`
	} `json:"data"`

	NextPageToken string `json:"nextPageToken"`
}
