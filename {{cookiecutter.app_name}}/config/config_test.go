package config

import (
{% if cookiecutter.use_ginkgo_testing == "y" %}
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
{% else %}
	"testing"
{% endif %}
)

{% if cookiecutter.use_ginkgo_testing == "y" %}
var _ = Describe("Config", func() {
	var cfg = defaultConfig
	BeforeEach(func() {
		//TODO do something
	})

	Describe("Example Parameters", func() {
		Context("with a string value set", func() {
			cfg.Set("foo", "bar")
			It("should be able to get it again", func() {
				Expect(cfg.Get("foo")).To(Equal("bar"))
			})
		})

		Context("with a typed value set", func() {
			cfg.Set("fooInt", 123)
			It("should be able to get it as typed value", func() {
				Expect(cfg.GetInt("fooInt")).To(Equal(123))
			})
		})
	})
})
{% else %}
func TestConfig(t *testing.T) {
	var cfg = defaultConfig

	want := "bar"
	cfg.Set("foo", want)
	got := cfg.Get("foo")
	if got != want {
		t.Errorf("Expected %s but got %s", want, got)
	}

	wantInt := 123
	cfg.Set("fooInt", wantInt)
	gotInt := cfg.Get("fooInt")
	if gotInt != wantInt {
		t.Errorf("Expected %d but got %d", wantInt, gotInt)
	}

}
{% endif %}
