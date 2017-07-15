package auth


import (
    "log"

    "github.com/dgrijalva/jwt-go"
)


type UserData struct {
    Email string `json:"email"`
}


func jwtSecretKeyFunc(secret string) jwt.Keyfunc {
    return func(t *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    }
}


func ReadToken(token, secret string) (UserData, error) {
    t, err := jwt.Parse(token, jwtSecretKeyFunc(secret))
    if err != nil {
        log.Println(err)
        return UserData{}, err
    }
    claims := t.Claims.(jwt.MapClaims)
    return UserData{claims["email"].(string)}, nil
}
