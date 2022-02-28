package config

import (
	{% if cookiecutter.use_ginkgo_testing == "y" %}
	"testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	{% endif %}
)

{% if cookiecutter.use_ginkgo_testing == "y" %}
func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}
{% else %}
// unused - please remove this file
{% endif %}
