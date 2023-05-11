package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Define a custom ingress struct
type MyIngress struct {
	Name     string   `yaml:"name"`
	Upstream string   `yaml:"upstream"`
	Rules    []string `yaml:"rules"`
}

func main() {
	// Example Apisix ingress YAML
	apisixYAML := `
name: my-ingress
upstream: my-upstream
rules:
- path: /foo
  backend: http://foo.com
- path: /bar
  backend: http://bar.com
`

	// Unmarshal the Apisix ingress YAML into our custom struct
	var myIngress MyIngress
	err := yaml.Unmarshal([]byte(apisixYAML), &myIngress)
	if err != nil {
		panic(err)
	}

	// Print the unmarshalled struct
	fmt.Printf("%+v\n", myIngress)
}

// type MyCustom struct{
// 	ApiVersion string `yaml: "api"`
// 	Name string `yaml: "name"`
// 	Metadata string `yaml: "metadata"`
// }
