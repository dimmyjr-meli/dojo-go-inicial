# Desafio 

Criar uma aplicação para receber os dados de cada host. Que deva ser incluída na regras de mapeamento do **NGIEX**.

Entrada dos dados:
+ comentario explicativo da regra, como: para qual aplicativo esta sendo cadastrado
+ location
+ se é publica
+ regras
    + host
    + scope

Depois de cadastrado as hosts, deve ser possível gerar o arquivo **api.conf**

**Exemplos de entrada de host**
+ comments: Tax Collected Source Gateway
+ location: " ~ ^/taxes-gateway "
+ http_x_public: true
+ regras:
  + regras 1 
    + scope: stage
    + host: http://stage.tax-collected-source-gateway.melifrontends.com;
  + regras 2
    + scope: dev
    + host http://dev.tax-collected-source-gateway.melifrontends.com;
  + regra padrao, sem scope
    + proxy_pass http://prod.tax-collected-source-gateway.melifrontends.com;



**Exemplos de saida com a regras cadastradas:**
```
#  Tax Collected Source Gateway
location ~ ^/taxes-gateway {
  if ($http_x_public = "true") { return 403; break;} # NOT PUBLIC API

  if ($http_x_api_scope = "stage") {
    proxy_pass http://stage.tax-collected-source-gateway.melifrontends.com;
    break;
  }

  if ($http_x_api_scope = "dev") {
    proxy_pass http://dev.tax-collected-source-gateway.melifrontends.com;
    break;
  }

  proxy_pass http://prod.tax-collected-source-gateway.melifrontends.com;
}

```
