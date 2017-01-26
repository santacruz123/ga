package ga

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	It("Response", func() {

		reader, err := os.Open("fixtures/response.json")
		Expect(err).To(Succeed())

		r := &HelperResponse{}

		json.NewDecoder(reader).Decode(r)

		data := r.Export()

		Expect(err).To(Succeed())
		Expect(data).NotTo(BeEmpty())
	})

	It("CSV", func() {

		reader, err := os.Open("fixtures/response.json")
		Expect(err).To(Succeed())

		r := &HelperResponse{}
		json.NewDecoder(reader).Decode(r)

		var csv bytes.Buffer

		Expect(r.CSV(&csv)).To(Succeed())

		fmt.Println(csv.String())

	})
})
