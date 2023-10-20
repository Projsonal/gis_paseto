package gis_paseto

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/petapedia/peda"
)

func init() {
	functions.HTTP("petapedia", petaSal2Pos)
}

func petaSal2Pos(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://jscroot.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://jscroot.github.io")
	fmt.Fprintf(w, peda.GCFPostHandler("PASETOPRIVATEKEY", "MONGOJAMBE", "petapedia", "user", r))

}

func GetToken(r *http.Request) string {
    return r.Header.Get("Authorization")
}