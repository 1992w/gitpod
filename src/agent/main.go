package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET(`/gitpod-core-dev/build/supervisor:commit-b4fc228990e4325ebf7b3a8079c41ac1437b1d2c`, func(context *gin.Context) {
		context.File("supervisor.html")
	})
	r.GET(`/v2/gitpod-core-dev/build/supervisor/manifests/commit-b4fc228990e4325ebf7b3a8079c41ac1437b1d2c`, func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"schemaVersion": 2,
			"mediaType":     "application/vnd.docker.distribution.manifest.v2+json",
			"config": map[string]interface{}{
				"mediaType": "application/vnd.docker.container.image.v1+json",
				"size":      3386,
				"digest":    "sha256:7f9b228096121c7ee34fbfc57a464e7827efa340423904cc4e89e2cb74ab6a60",
			},
			"layers": []map[string]interface{}{
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      244772,
					"digest":    "sha256:6846ff1102b50d41907585fcdaf1f359e6294765b03bbd8b3559855b0e4cf261",
				},
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      12059162,
					"digest":    "sha256:6ed688a83490707c037a2e8375703ef023ed73fdd1fb71dc1c338c1464a39aa5",
				},
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      152,
					"digest":    "sha256:85d305c2f7f9b43113f6ff61547846c52efe995ea79077772a2f522c22168488",
				},
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      1507113,
					"digest":    "sha256:e07526d3d0fde68a5fabc6d83f054ec8f5643140461dbb35e56912fb513ba4b1",
				},
				{
					"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
					"size":      1264047,
					"digest":    "sha256:8e687ae5c7cd2265540c6d07db4ec3d6cec18e717025177ff1cdcfba38230a1c",
				},
			},
		})
	})
	r.Run(":80")
}
