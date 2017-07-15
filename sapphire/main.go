package main


import (
    "encoding/json"
    "net/http"
    "net/http/httputil"

    "github.com/projectweekend/sapphire/auth"
    "github.com/projectweekend/sapphire/config"
)


func main()  {
    conf := config.Options()

    proxyHandler := httputil.NewSingleHostReverseProxy(conf.DstURL)

    jwt := jwtMiddleware(conf.JWTSecret)
    http.Handle("/", jwt(proxyHandler))
    http.ListenAndServe(conf.Host, nil)
}


type middleware func(http.Handler) http.Handler


func jwtMiddleware(jwtSecret string) middleware {
    return func(next http.Handler) http.Handler {
        handleAuth := func(w http.ResponseWriter, r *http.Request)  {
            tokenVal := r.Header.Get("Authorization")
            if tokenVal == "" {
                http.Error(w, http.StatusText(401), 401)
                return
            }
            user, err := auth.ReadToken(tokenVal, jwtSecret)
            if err != nil {
                http.Error(w, http.StatusText(401), 401)
                return
            }
            sapphireUser, _ := json.Marshal(user)
            w.Header().Set("X-Sapphire-User", string(sapphireUser))
            next.ServeHTTP(w, r)
        }
        return http.HandlerFunc(handleAuth)
    }
}
