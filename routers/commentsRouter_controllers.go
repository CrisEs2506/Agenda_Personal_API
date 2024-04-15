package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:AgendaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:ContactoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agenda_personal/controllers:NumeroTelefonoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
