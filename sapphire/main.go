package main

import (
    "log"

    "github.com/projectweekend/sapphire/config"
)

func main()  {
    conf := config.Options()
    log.Println(conf)
}
