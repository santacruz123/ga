package ga_test

import (
	"bytes"
	"os"
	"strings"

	"github.com/santacruz123/ga"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("External", func() {
	Context("External", func() {
		It("GA", func() {
			req := ga.New(os.Getenv("ACCESS"))

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
	})
})
