# juridigo_api_pagamentos
Servidor


# Obtenção e inicialização do projeto

```
cd go/src/github.com/
mkdir juridigo
cd juridigo
git clone https://github.com/andre250/juridigo_api_pagamentos.git
cd juridigo_api_pagamentos/
dep ensure 
gin -i main.go
```

# Endpoints

## Pagamentos
```
http://.../pagamento

METHOD: POST

Descrição: Registra um pagamento

Body: 
{
	"trabalhoId": [String],
	"usuarioId":[String],
	"valor":[Float64],
	"dataConclusao":[String | unix timestamp]
}
```
```
http://.../pagamento?pagamento={id}

METHOD: PUT

Descrição: Registra um pagamento

Parametros:
- id = string identificador do pagamento atualizado

Body: 
{
	"campo":"valor"
}
```
```
http://.../pagamento?trabalho={id}

METHOD: GET

Descrição: Busca todos os pagamentos de um trabalho

Parametros:
- id = string identificador do trabalho 

```

```
http://.../pagamento?usuario={id}&status={status}

METHOD: GET

Descrição: Busca todos os pagamentos de um usuario por status

Parametros:
- id = string identificador do usuario
- status = booleano [true|false]   

```