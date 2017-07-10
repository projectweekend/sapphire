package main


import (
    "log"
    "net/http"
    "net/http/httputil"

    "github.com/projectweekend/sapphire/config"
)


func main()  {
    conf := config.Options()

    proxyHandler := httputil.NewSingleHostReverseProxy(conf.DstURL)

    http.Handle("/", jwtMiddleware(proxyHandler))
    http.ListenAndServe(":9009", nil)

}


func jwtMiddleware(next http.Handler) http.Handler {
    handleAuth := func(w http.ResponseWriter, r *http.Request)  {
        authorization := r.Header.Get("Authorization")
        if authorization == "" {
            http.Error(w, http.StatusText(401), 401)
            return
        }
        // TODO: validate and decode JWT token from header
        log.Println(authorization)
        next.ServeHTTP(w, r)
    }
    return http.HandlerFunc(handleAuth)
}
