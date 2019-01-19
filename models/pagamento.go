package models

import (
	"gopkg.in/mgo.v2/bson"
)

/*
Pagamento - Modelo estrutural de pagamento
*/
type Pagamento struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	TrabalhoID    string        `bson:"trabalhoId,omitempty" json:"trabalhoId,omitempty"`
	UsuarioID     string        `bson:"usuarioId,omitempty" json:"usuarioID,omitempty"`
	Valor         float64       `bson:"valor,omitempty" json:"valor,omitempty"`
	DataConclusao string        `bson:"dataConclusao,omitempty" json:"dataConclusao,omitempty"`
	DataPagamento string        `bson:"dataPagamento,omitempty" json:"dataPagamento,omitempty"`
	Status        bool          `bson:"status" json:"status"`
}
