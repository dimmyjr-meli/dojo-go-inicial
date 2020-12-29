package main

import (
	"fmt"
)

func main() {
	regras := []Regras{
		Regras{
			scope: "stage",
			host:  "http://stage.tax-collected-source-gateway.melifrontends.com",
		},
		Regras{
			scope: "dev",
			host:  "http://dev.tax-collected-source-gateway.melifrontends.com",
		},
		Regras{
			host: "http://prod.tax-collected-source-gateway.melifrontends.com",
		},
	}
	Entrada ("Tax Collected Source Gateway"," ~ ^/taxes-gateway ",false, regras)
}


func Hello(s string) string {
	return "Hello " + s
}

 type Regras struct {
	scope string
	host string
 }

func Entrada(comments string, location string, http_x_public bool, listRegras []Regras) {
	fmt.Println("#   ", comments)
	fmt.Printf("location %s {\n", location)

	if (!http_x_public) {
		fmt.Println("  if ($http_x_public = \"true\") { return 403; break;} # NOT PUBLIC API\n")
	}
	for _, val := range listRegras{
		if val.scope != "" {
			fmt.Printf("  if ($http_x_api_scope = \"%s\") {\n", val.scope)
			fmt.Printf("    proxy pass %s; \n    break;\n", val.host)
			fmt.Println("  }\n")
		} else {
				fmt.Printf("  proxy pass %s; \n", val.host)
		}
	}
	fmt.Println("}")



}