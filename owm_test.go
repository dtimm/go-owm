package owm_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/dtimm/go-owm"
)

func TestCurrentWeather(t *testing.T) {
	t.Run("it calls the weather endpoint with an API key", func(t *testing.T) {
		g := NewGomegaWithT(t)

		do := func(r http.Request) (*http.Response, error) {
			q := r.URL.Query()
			g.Expect(q["appid"]).To(Equal([]string{"test-api-key"}))
			return nil, nil
		}

		w := owm.New(do, "test-api-key")
		_, err := w.Current("Denver")
		g.Expect(err).ToNot(HaveOccurred())
	})
}
