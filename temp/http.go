package temp

import "net/http"

func FormValue(r *http.Request, key string) string {
	return r.FormValue(key)
}
