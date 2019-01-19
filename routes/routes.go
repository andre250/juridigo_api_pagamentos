package routes

import (
	"github.com/juridigo/juridigo_api_pagamentos/controllers"

	"github.com/juridigo/juridigo_api_pagamentos/helpers"
	"github.com/juridigo/juridigo_api_pagamentos/models"
)

/*
Routes - Controlador de rotas do microsserviço
*/
func Routes() {
	helpers.APIDisperser("/pagamento",
		models.DefaultAPI{SubPath: "", Handler: controllers.PaymentDisperser, Auth: false},
	)
}
