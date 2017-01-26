package ga

import "os"

var access = os.Getenv("ACCESS")

type metricType string

const (
	metricTypeUnspecified metricType = "METRIC_TYPE_UNSPECIFIED"
	metricTypeInteger     metricType = "INTEGER"
	metricTypeFloat       metricType = "FLOAT"
	metricTypeCurrency    metricType = "CURRENCY"
	metricTypePercent     metricType = "PERCENT"
	metricTypeTime        metricType = "TIME"
)

type dateRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type dimension struct {
	Name             string   `json:"name"`
	HistogramBuckets []string `json:"histogramBuckets,omitempty"`
}

type metric struct {
	Expression     string     `json:"expression"`
	Alias          string     `json:"alias,omitempty"`
	FormattingType metricType `json:"formattingType,omitempty"`
}

type request struct {
	ViewID     string      `json:"viewId"`
	DateRange  []dateRange `json:"dateRanges"`
	Dimensions []dimension `json:"dimensions,omitempty"`
	Metrics    []metric    `json:"metrics,omitempty"`
	PageToken  string      `json:"pageToken,omitempty"`
	PageSize   int64       `json:"pageSize,omitempty"`
}

type response struct {
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
