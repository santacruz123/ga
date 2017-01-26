package ga

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//HelperRequest struct
type HelperRequest struct {
	accessCode string
	Request    Request
}

//New request constructor
func New(accessCode string) *HelperRequest {
	return &HelperRequest{
		accessCode: accessCode,
		Request: Request{
			PageSize: 10000,
		},
	}
}

//Dimension add
func (h *HelperRequest) Dimension(dim string) {
	h.Request.Dimensions = append(
		h.Request.Dimensions,
		Dimension{
			Name: dim,
		},
	)
}

//Metric add
func (h *HelperRequest) Metric(expr, alias string, style MetricType) {

	m := Metric{}
	m.Expression = expr

	if alias != "" {
		m.Alias = alias
	}

	if style != "" {
		m.FormattingType = style
	}

	h.Request.Metrics = append(h.Request.Metrics, m)
}

//PageSize set
func (h *HelperRequest) PageSize(size int64) {
	h.Request.PageSize = size
}

//ViewID set
func (h *HelperRequest) ViewID(viewID string) {
	h.Request.ViewID = viewID
}

//DateRange add
func (h *HelperRequest) DateRange(start, end string) {
	h.Request.DateRange = append(
		h.Request.DateRange,
		DateRange{
			StartDate: start,
			EndDate:   end,
		},
	)
}

//Do request
func (h *HelperRequest) Do() (res *HelperResponse, err error) {

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

	tmpHelpResp := HelperResponse{}

	if err = json.NewDecoder(resHTTP.Body).Decode(&tmpHelpResp); err != nil {
		return
	}

	res = &tmpHelpResp

	return
}

func (h *HelperRequest) marshal() ([]byte, error) {
	tmpStruc := struct {
		ReportRequests []Request `json:"reportRequests"`
	}{
		ReportRequests: []Request{h.Request},
	}

	return json.Marshal(&tmpStruc)
}
