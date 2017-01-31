package ga

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//Request struct
type Request struct {
	accessCode string
	request    request
}

//New constructor
func New(accessCode string) *Request {
	return &Request{
		accessCode: accessCode,
		request: request{
			PageSize: 10000,
		},
	}
}

//Dimension add
func (h *Request) Dimension(dim string) {
	h.request.Dimensions = append(
		h.request.Dimensions,
		dimension{
			Name: dim,
		},
	)
}

//Metric add
func (h *Request) Metric(expr, alias, style string) {

	m := metric{}
	m.Expression = expr

	if alias != "" {
		m.Alias = alias
	}

	if style != "" {
		m.FormattingType = metricType(style)
	}

	h.request.Metrics = append(h.request.Metrics, m)
}

//PageSize setter
func (h *Request) PageSize(size int64) {
	h.request.PageSize = size
}

//ViewID setter
func (h *Request) ViewID(viewID string) {
	h.request.ViewID = viewID
}

//DateRange add
func (h *Request) DateRange(start, end string) {
	h.request.DateRange = append(
		h.request.DateRange,
		dateRange{
			StartDate: start,
			EndDate:   end,
		},
	)
}

//Do request
func (h *Request) Do() (res *Response, err error) {

	payload, err := h.marshal()
	if err != nil {
		return
	}

	req, err := http.NewRequest(
		"POST",
		"https://analyticsreporting.googleapis.com/v4/reports:batchGet",
		bytes.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+h.accessCode)

	resHTTP, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if resHTTP.StatusCode != 200 {
		err = errors.New(resHTTP.Status)
		return
	}

	defer resHTTP.Body.Close()

	tmpHelpResp := Response{}

	if err = json.NewDecoder(resHTTP.Body).Decode(&tmpHelpResp); err != nil {
		return
	}

	res = &tmpHelpResp

	return
}

func (h *Request) marshal() ([]byte, error) {
	tmpStruc := struct {
		ReportRequests []request `json:"reportRequests"`
	}{
		ReportRequests: []request{h.request},
	}

	return json.Marshal(&tmpStruc)
}
