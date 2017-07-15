package main


import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
)

func main()  {
    jwtSecret := os.Getenv("JWT_SECRET")
    token, err := makeToken(userData{"test@test.com"}, jwtSecret)
    if err != nil {
        log.Fatal("Failed to create JWT")
    }

    client := &http.Client{}

    req, err := http.NewRequest("GET", "http://sapphire:9009/whatever", nil)
    if err != nil {
        log.Fatal("Failed to create request")
    }
    req.Header.Add("Authorization", token)

    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Failed to do GET request")
    }
    log.Println(resp)
}


type customClaims struct {
    userData
    jwt.StandardClaims
}


type userData struct {
    Email string `json:"email"`
}


func makeToken(ud userData, secret string) (string, error) {
    claims := customClaims{
        ud,
        jwt.StandardClaims{
            ExpiresAt: int64(time.Now().Unix() + 1000),
            Issuer:    "test",
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
