module github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}

require (
	{% if cookiecutter.use_logrus_logging == "y" -%}github.com/sirupsen/logrus v1.4.1{%- endif %}
	{% if cookiecutter.use_cobra_cmd == "y" -%}github.com/spf13/cobra v0.0.3{%- endif %}
	{% if cookiecutter.use_viper_config == "y" -%}github.com/spf13/viper v1.3.2{%- endif %}
	{% if cookiecutter.use_ginkgo_testing == "y" -%}
	github.com/onsi/ginkgo/v2 v2.1.3
	github.com/onsi/gomega v1.18.1
	{%- endif %}
)
