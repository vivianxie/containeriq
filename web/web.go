package web

import (
	"bitbucket.org/containeriq/wharf/router"
	"gopkg.in/macaron.v1"
)

// "bitbucket.org/containeriq/wharf/middleware"
// "bitbucket.org/containeriq/wharf/router"

func SetWharfMacaron(m *macaron.Macaron) {
	//Setting Database

	// if err := middleware.Initfunc(); err != nil {
	// 	fmt.Printf("Init middleware error %s", err.Error())
	// }
	//
	// //Setting Middleware
	// middleware.SetMiddlewares(m)

	// m.Use(macaron.Statics(macaron.StaticOptions{}, "static"))
	m.Use(macaron.Static("static"))

	//Setting Router
	router.SetRouters(m)

}
