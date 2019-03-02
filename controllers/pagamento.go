package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/juridigo/juridigo_api_pagamentos/utils"

	"github.com/juridigo/juridigo_api_pagamentos/models"

	"github.com/juridigo/juridigo_api_pagamentos/helpers"
)

/*
PaymentDisperser - Controlador de chamadas para Api de pagamentos
*/
func PaymentDisperser(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "POST", "PUT", "GET") != nil {
		return
	}

	if r.Method == "POST" {
		createPayment(w, r)
	} else if r.Method == "PUT" {
		updatePayment(w, r)
	} else if r.Method == "GET" {
		getPayment(w, r)
	}
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	var pagamento models.Pagamento
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&pagamento)

	pagamento.Status = "0"

	err := helpers.Db().Update("propostas", bson.M{"_id": bson.ObjectIdHex(pagamento.PropostaID)}, bson.M{"$set": bson.M{
		"status": "4",
	}})

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte(`{"msg":"Erro Interno"}`))
		return
	}

	err = helpers.Db().Insert("pagamentos", &pagamento)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg":"parametros inválidos na chamada"}`))
		return
	}

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Pagamento criado com sucesso!"}`))
}

func updatePayment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("pagamento")

	var pagamento models.Pagamento
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&pagamento)

	if pagamento.Status == "1" {
		pagamento.DataPagamento = time.Now().UnixNano() / int64(time.Millisecond)
	}
	err := helpers.Db().Update("pagamentos", bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": pagamento})
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg":"Identificador inválido na chamada"}`))
		return
	}

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Pagamento atualizado com sucesso!"}`))
}

func getPayment(w http.ResponseWriter, r *http.Request) {
	trabalho := r.URL.Query().Get("trabalho")
	usuario := r.URL.Query().Get("usuario")
	status := r.URL.Query().Get("status")

	var finalStatus bool
	if status != "" {
		if status == "true" {
			finalStatus = true
		} else {
			finalStatus = false
		}
	}
	var err error
	var itens []interface{}
	if trabalho != "" {
		itens, err = helpers.Db().Find("pagamentos", bson.M{"trabalhoId": trabalho}, -1)
	} else if usuario != "" {
		if status != "" {
			itens, err = helpers.Db().Find("pagamentos", bson.M{"usuarioId": usuario, "status": finalStatus}, -1)
		} else {
			itens, err = helpers.Db().Find("pagamentos", bson.M{"usuarioId": usuario}, -1)
		}
	} else if status != "" {
		fmt.Println("oi")
		itens, err = helpers.Db().Find("pagamentos", bson.M{"status": finalStatus}, -1)
	} else {
		itens, err = helpers.Db().Find("pagamentos", bson.M{}, -1)
	}

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}

	listItens, _ := bson.MarshalJSON(itens)
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write(listItens)
}
