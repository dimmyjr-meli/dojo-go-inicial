package main

import "fmt"

func main() {

	var regra = []Regra {
		Regra{ "http://stage.tax-collected-source-gateway.melifrontends.com", "stage"},
		Regra{ "http://dev.tax-collected-source-gateway.melifrontends.com", "dev"},
	}
	var rules = Rules{"Tax Collected Source Gateway", "location ~ ^/taxes-gateway ", true, "http://prod.tax-collected-source-gateway.melifrontends.com", regra}

	Output(rules)
}

type Rules struct {
	comment string
	location string
	http_x_public bool
	proxy_pass string
	regras []Regra

}

type Regra struct {
	host  string
	scope string
}

func Output(rules Rules){
	fmt.Printf("# %s  \n", rules.comment)
	fmt.Printf("%s {\n", rules.location)

	if rules.http_x_public {
		fmt.Printf("if(http_x_public = \"%t\") { return 403; \t break;} # NOT PUBLIC API \n", rules.http_x_public)
	}

	for i := 0; i < len(rules.regras); i++ {
		fmt.Printf("if ($http_x_api_scope = \"%s\"){ \n proxy_pass %s; \n\tbreak; \n } \n\n", rules.regras[i].scope, rules.regras[i].host)
	}

	fmt.Printf("proxy_pass %s; \n}", rules.proxy_pass)


}