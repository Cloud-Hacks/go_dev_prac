package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/apache/apisix-ingress-controller/pkg/types/apisix/v1"
	_ "gopkg.in/yaml.v2"
)

func main() {
	// Read custom ingress YAML file
	data, err := ioutil.ReadFile("custom-ingress.yaml")
	if err != nil {
		panic(err)
	}

	// Unmarshal custom ingress YAML into a struct
	var customIngress struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Name string `yaml:"name"`
		} `yaml:"metadata"`
		Spec struct {
			Rules []struct {
				Host string `yaml:"host"`
				HTTP struct {
					Paths []struct {
						Path    string `yaml:"path"`
						Backend struct {
							ServiceName string `yaml:"serviceName"`
							ServicePort int    `yaml:"servicePort"`
						} `yaml:"backend"`
					} `yaml:"paths"`
				} `yaml:"http"`
			} `yaml:"rules"`
		} `yaml:"spec"`
	}

	err = yaml.Unmarshal(data, &customIngress)
	if err != nil {
		panic(err)
	}

	// Validate custom ingress YAML against Apisix ingress YAML
	var apisixIngress v1.ApisixRoute
	apisixIngress.Hosts = append(apisixIngress.Hosts, customIngress.Spec.Rules[0].Host)
	apisixIngress.Paths = []v1.ApisixRoutePath{
		{
			Path: customIngress.Spec.Rules[0].HTTP.Paths[0].Path,
			Backend: v1.ApisixRouteBackend{
				ServiceName: customIngress.Spec.Rules[0].HTTP.Paths[0].Backend.ServiceName,
				ServicePort: customIngress.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort,
			},
		},
	}

	// Marshal Apisix ingress YAML to JSON
	apisixIngressJSON, err := json.Marshal(apisixIngress)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(apisixIngressJSON))
}
