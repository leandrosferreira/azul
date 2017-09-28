package main

import "fmt"
import _ "github.com/lsferreira42/azul/api"
import _ "github.com/lsferreira42/azul/backends"
import _ "github.com/lsferreira42/azul/caching"
import _ "github.com/lsferreira42/azul/dns"
import _ "github.com/lsferreira42/azul/resolver"

func main() {
	fmt.Println("Hello from Azul!")
}
