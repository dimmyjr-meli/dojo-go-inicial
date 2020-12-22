package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	result := createHandler("^/taxes-gateway", "# Tax Collected Source Gateway",[] string {
		createRule("", "", true),
		createRule("http://stage.taxes-collected-source-gateway.melifrontends.com", "stage", false),
		createRule("http://dev.taxes-collected-source-gateway.melifrontends.com", "dev", false),
		createRule("http://prod.taxes-collected-source-gateway.melifrontends.com", "", false),
	})
	fmt.Printf(result)
	err := ioutil.WriteFile("rules.conf", []byte(result), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func createHandler(path, commentBlock string, rules[] string) string {
	proxyRulesString := commentBlock + "\n"
	proxyRulesString += "location ~ "+ path + " {\n\t"
	for i := 0; i < len(rules); i++ {
		proxyRulesString += rules[i] + "\n\t"
	}
	proxyRulesString += "\n}"
	return proxyRulesString
}

func createRule(host, scope string, public bool) string {
	if public == true{
		return "if ($http_x_public = \"true\") { return 403; break;} # NOT PUBLIC API \n"
	}

	if len(scope) > 0 {
		return "if ($http_x_api_scope = \"" + scope + "\") {\n\t\tproxy_pass " + host + ";\n\t\tbreak;\n\t}\n"
	}
	return "proxy_pass " + host + ";"
}



