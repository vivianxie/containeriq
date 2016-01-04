package router

import (
	"bitbucket.org/containeriq/wharf/handler"
	"gopkg.in/macaron.v1"
)

func SetRouters(m *macaron.Macaron) {
	//Docker Registry & Hub V1 API
	m.Group("/v1", func() {
		m.Get("/_ping", handler.GetPingV1Handler)
	})
	// m.Get("/_ping", handler.GetPingV1Handler)
	//
	// m.Get("/users", handler.GetUsersV1Handler)
	// m.Post("/users", handler.PostUsersV1Handler)

	// m.Group("/repositories", func() {
	// 	m.Put("/:namespace/:repository/tags/:tag", handler.PutTagV1Handler)
	// 	m.Put("/:namespace/:repository/images", handler.PutRepositoryImagesV1Handler)
	// 	m.Get("/:namespace/:repository/images", handler.GetRepositoryImagesV1Handler)
	// 	m.Get("/:namespace/:repository/tags", handler.GetTagV1Handler)
	// 	m.Put("/:namespace/:repository", handler.PutRepositoryV1Handler)
	// })

	// 	m.Group("/images", func() {
	// 		m.Get("/:imageId/ancestry", handler.GetImageAncestryV1Handler)
	// 		m.Get("/:imageId/json", handler.GetImageJSONV1Handler)
	// 		m.Get("/:imageId/layer", handler.GetImageLayerV1Handler)
	// 		m.Put("/:imageId/json", handler.PutImageJSONV1Handler)
	// 		m.Put("/:imageId/layer", handler.PutImageLayerv1Handler)
	// 		m.Put("/:imageId/checksum", handler.PutImageChecksumV1Handler)
	// 	})

}
