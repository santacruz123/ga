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
	request    request
}

//New constructor
func New(accessCode string) *HelperRequest {
	return &HelperRequest{
		accessCode: accessCode,
		request: request{
			PageSize: 10000,
		},
	}
}

//Dimension add
func (h *HelperRequest) Dimension(dim string) {
	h.request.Dimensions = append(
		h.request.Dimensions,
		dimension{
			Name: dim,
		},
	)
}

//Metric add
func (h *HelperRequest) Metric(expr, alias, style string) {

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

//PageSize set
func (h *HelperRequest) PageSize(size int64) {
	h.request.PageSize = size
}

//ViewID set
func (h *HelperRequest) ViewID(viewID string) {
	h.request.ViewID = viewID
}

//DateRange add
func (h *HelperRequest) DateRange(start, end string) {
	h.request.DateRange = append(
		h.request.DateRange,
		dateRange{
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
		ReportRequests []request `json:"reportRequests"`
	}{
		ReportRequests: []request{h.request},
	}

	return json.Marshal(&tmpStruc)
}
