package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["park/controllers:ParkingSpotController"] = append(beego.GlobalControllerRouter["park/controllers:ParkingSpotController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:ParkingSpotController"] = append(beego.GlobalControllerRouter["park/controllers:ParkingSpotController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:ParkingSpotController"] = append(beego.GlobalControllerRouter["park/controllers:ParkingSpotController"],
        beego.ControllerComments{
            Method: "DeleteList",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:ParkingSpotController"] = append(beego.GlobalControllerRouter["park/controllers:ParkingSpotController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:ParkingSpotController"] = append(beego.GlobalControllerRouter["park/controllers:ParkingSpotController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:ParkingSpotController"] = append(beego.GlobalControllerRouter["park/controllers:ParkingSpotController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:UserController"] = append(beego.GlobalControllerRouter["park/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:UserController"] = append(beego.GlobalControllerRouter["park/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:UserController"] = append(beego.GlobalControllerRouter["park/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteList",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:UserController"] = append(beego.GlobalControllerRouter["park/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:UserController"] = append(beego.GlobalControllerRouter["park/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:UserController"] = append(beego.GlobalControllerRouter["park/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:WebSocketController"] = append(beego.GlobalControllerRouter["park/controllers:WebSocketController"],
        beego.ControllerComments{
            Method: "Information",
            Router: `/information`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["park/controllers:WebSocketController"] = append(beego.GlobalControllerRouter["park/controllers:WebSocketController"],
        beego.ControllerComments{
            Method: "Logcat",
            Router: `/logcat`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
