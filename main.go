package function

import (
	"net/http"
	"os"

	src "github.com/leafy-ai/backend_b2/src"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gin-gonic/gin"
)

func init() {
	entryPoint := os.Getenv("FUNCTION_TARGET")
	if entryPoint == "" {
		entryPoint = "BlogsAPI"
	}
	functions.HTTP(entryPoint, EntryPoint)
}

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	g := gin.Default()
	src.GetRoutes(g)
	g.Handler().ServeHTTP(w, r)
}
