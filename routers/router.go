// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/agenda_personal/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/agenda",
			beego.NSInclude(
				&controllers.AgendaController{},
			),
		),

		beego.NSNamespace("/contacto",
			beego.NSInclude(
				&controllers.ContactoController{},
			),
		),

		beego.NSNamespace("/numero_telefono",
			beego.NSInclude(
				&controllers.NumeroTelefonoController{},
			),
		),

		beego.NSNamespace("/correo_electronico",
			beego.NSInclude(
				&controllers.CorreoElectronicoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
