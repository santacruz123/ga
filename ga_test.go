package ga

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	Context("Request", func() {

		It("Should generate", func() {

			req := New("123")

			req.ViewID("112236938")
			req.DateRange("2017-01-01", "2017-01-02")

			req.Dimension("ga:deviceCategory")
			req.Dimension("ga:campaign")
			req.Dimension("ga:adGroup")

			req.Metric("ga:users", "", "")
			req.Metric("ga:sessions", "", "")

			reqBytes, err := ioutil.ReadFile("fixtures/request.json")
			Expect(err).To(Succeed())

			reqString := string(reqBytes)
			reqString = strings.Replace(reqString, " ", "", -1)
			reqString = strings.Replace(reqString, "\n", "", -1)

			byt, err := req.marshal()
			Expect(err).To(Succeed())

			Expect(string(byt)).To(Equal(reqString))
		})
	})

	It("GA request", func() {

		req := New(os.Getenv("ACCESS"))

		req.ViewID("112236938")
		req.DateRange("2017-01-01", "2017-01-02")

		req.Dimension("ga:deviceCategory")
		req.Dimension("ga:campaign")
		req.Dimension("ga:adGroup")

		req.Metric("ga:users", "", "")
		req.Metric("ga:sessions", "", "")

		res, err := req.Do()
		Expect(err).To(Succeed())

		var csv bytes.Buffer

		Expect(res.CSV(&csv)).To(Succeed())
		strings.Contains(csv.String(), "1259490_LBX_160505_10moneybackbonus_core")
	})

	It("Response", func() {

		reader, err := os.Open("fixtures/response.json")
		Expect(err).To(Succeed())

		r := &Response{}

		json.NewDecoder(reader).Decode(r)

		data := r.Export()

		Expect(err).To(Succeed())
		Expect(data).NotTo(BeEmpty())
	})

	It("CSV", func() {

		reader, err := os.Open("fixtures/response.json")
		Expect(err).To(Succeed())

		r := &Response{}
		json.NewDecoder(reader).Decode(r)

		var csv bytes.Buffer

		Expect(r.CSV(&csv)).To(Succeed())

		strings.Contains(csv.String(), "1259490_LBX_160505_10moneybackbonus_core")

	})
})
