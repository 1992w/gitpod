package main

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	r := gin.Default()
	r.GET("/gitpod-core-dev/build/supervisor:commit-96916c5a955b991a94b23aabe5cdbd9fd7881dba", func(context *gin.Context) {
		context.File("supervisor.html")
	})
	r.HEAD(`/v2/gitpod-core-dev/build/supervisor/manifests/commit-96916c5a955b991a94b23aabe5cdbd9fd7881dba`, func(context *gin.Context) {
		w := context.Writer
		w.Header().Add("Content-Type", "application/vnd.docker.distribution.manifest.v2+json")
		w.Header().Add("Docker-Distribution-API-Version", "registry/2.0")
		w.Header().Add("Docker-Content-Digest", "sha256:316572d4c8d0e0e365b116d4e336f4086a12b66cb1888d6cb487bc9eebb6f7d0")
		w.Header().Add("X-XSS-Protection", "0")
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		w.WriteHeader(200)
	})

	r.GET(`/v2/gitpod-core-dev/build/supervisor/manifests/commit-96916c5a955b991a94b23aabe5cdbd9fd7881dba`, func(context *gin.Context) {
		w := context.Writer
		w.Header().Add("Content-Type", "application/vnd.docker.distribution.manifest.v2+json")
		w.Header().Add("Docker-Distribution-API-Version", "registry/2.0")
		w.Header().Add("Docker-Content-Digest", "sha256:316572d4c8d0e0e365b116d4e336f4086a12b66cb1888d6cb487bc9eebb6f7d0")
		w.Header().Add("X-XSS-Protection", "0")
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		rsp, _ := json.Marshal(map[string]interface{}{
			"schemaVersion": 2,
			"mediaType":     "application/vnd.docker.distribution.manifest.v2+json",
			"config": map[string]interface{}{
				"mediaType": "application/vnd.docker.container.image.v1+json",
				"size":      3357,
				"digest":    "sha256:46d0cd90e651e8702c6f54c8c04da4bf8fed2edd335fdc1c4671f3946d4ead4f",
			},
			"layers": []map[string]interface{}{
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      245359,
					"digest":    "sha256:1e85f8ced572b47d9e8f7f8d837d577f5d142ea274e980e89caec3e0b047546f",
				},
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      12051307,
					"digest":    "sha256:2fdb575da91feb55c1265f88df0b080e20eba1d979fdf73d52615a3859904b75",
				},
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      153,
					"digest":    "sha256:c2b4524fd62f988eab62e87dc08687bdc9bf8f50a4ccc5dbb4c98f50cdac1792",
				},
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      1507111,
					"digest":    "sha256:c3b025001c78877c9456a7162ed3078961841b1501e5023dc1dbc353c57fd512",
				},
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      1264047,
					"digest":    "sha256:7a537eabd11811d0abe43b0c1b40ce69af58b2f923b5502fc6ea6bf623643791",
				},
			},
		})

		w.WriteHeader(200)
		w.Write(rsp)
	})

	r.GET("/v2/gitpod-core-dev/build/supervisor/manifests/sha256:2591c779eea183f98916c5064338459cd3f2b091c36ae161d983ac9ab8d4cf47", func(context *gin.Context) {
		w := context.Writer
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Docker-Distribution-API-Version", "registry/2.0")

		w.WriteHeader(404)
		w.WriteString(`{"errors":[{"code":"MANIFEST_UNKNOWN","message":"Failed to fetch \"sha256:2591c779eea183f98916c5064338459cd3f2b091c36ae161d983ac9ab8d4cf47\" from request \"/v2/gitpod-core-dev/build/supervisor/manifests/sha256:2591c779eea183f98916c5064338459cd3f2b091c36ae161d983ac9ab8d4cf47\"."}]}`)
	})
	// r.GET(`/gitpod-core-dev/build/supervisor:commit-b4fc228990e4325ebf7b3a8079c41ac1437b1d2c`, func(context *gin.Context) {
	// 	context.File("supervisor.html")
	// })
	// r.HEAD(`/v2/gitpod-core-dev/build/supervisor/manifests/commit-b4fc228990e4325ebf7b3a8079c41ac1437b1d2c`, func(context *gin.Context) {
	// 	context.JSON(200, nil)
	// })
	// r.GET(`/v2/gitpod-core-dev/build/supervisor/manifests/commit-b4fc228990e4325ebf7b3a8079c41ac1437b1d2c`, func(context *gin.Context) {
	// 	w := context.Writer
	// 	w.Header().Add("Content-Type", "application/vnd.docker.distribution.manifest.v2+json")
	// 	w.Header().Add("Docker-Distribution-API-Version", "registry/2.0")
	// 	w.Header().Add("Docker-Content-Digest", "sha256:04fc79fadf113024253ad379f1824c4f8aff46af8a235f227489941b7194d280")
	// 	rsp, _ := json.Marshal(map[string]interface{}{
	// 		"schemaVersion": 2,
	// 		"mediaType":     "application/vnd.docker.distribution.manifest.v2+json",
	// 		"config": map[string]interface{}{
	// 			"mediaType": "application/vnd.docker.container.image.v1+json",
	// 			"size":      3386,
	// 			"digest":    "sha256:7f9b228096121c7ee34fbfc57a464e7827efa340423904cc4e89e2cb74ab6a60",
	// 		},
	// 		"layers": []map[string]interface{}{
	// 			{
	// 				"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
	// 				"size":      244772,
	// 				"digest":    "sha256:6846ff1102b50d41907585fcdaf1f359e6294765b03bbd8b3559855b0e4cf261",
	// 			},
	// 			{
	// 				"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
	// 				"size":      12059162,
	// 				"digest":    "sha256:6ed688a83490707c037a2e8375703ef023ed73fdd1fb71dc1c338c1464a39aa5",
	// 			},
	// 			{
	// 				"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
	// 				"size":      152,
	// 				"digest":    "sha256:85d305c2f7f9b43113f6ff61547846c52efe995ea79077772a2f522c22168488",
	// 			},
	// 			{
	// 				"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
	// 				"size":      1507113,
	// 				"digest":    "sha256:e07526d3d0fde68a5fabc6d83f054ec8f5643140461dbb35e56912fb513ba4b1",
	// 			},
	// 			{
	// 				"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
	// 				"size":      1264047,
	// 				"digest":    "sha256:8e687ae5c7cd2265540c6d07db4ec3d6cec18e717025177ff1cdcfba38230a1c",
	// 			},
	// 		},
	// 	})

	// 	w.WriteHeader(200)
	// 	w.Write(rsp)

	// })
	r.Use(TlsHandler(443))
	r.RunTLS(":443", "test.pem", "test.key")
}

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
