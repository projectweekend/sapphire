package config


import (
    "flag"
    "log"
    "net/url"
)


type configOptions struct {
    DstURL *url.URL
    JWTSecret string
}


func Options() configOptions {
    dstURLstr := flag.String("destination_url", "", "URL for remote service where traffic will be forwarded")
    jwtSecret := flag.String("jwt_secret", "", "Secret used to validate JWTs")

    flag.Parse()

    if *dstURLstr == "" {
        log.Fatal("destination_url is required")
    }
    dstURL, err := url.Parse(*dstURLstr)
    if err != nil {
        log.Fatal("destination_url must be a valid URL")
    }

    if *jwtSecret == "" {
        log.Fatal("jwt_secret is required")
    }

    return configOptions{dstURL, *jwtSecret}
}
