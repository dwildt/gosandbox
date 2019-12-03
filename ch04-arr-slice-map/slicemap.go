package main

import "fmt"

var properties = make(map[string]string)

func activateFeature(key string) map[string]string {
	properties[key] = "on"
	return properties
}

func shutdownFeature(key string) map[string]string {
	properties[key] = "off"
	return properties
}

func isActive(key string) bool {
	return (properties[key] == "on")
}

func main() {

	activateFeature("login.google")
	activateFeature("login.facebook")

	if properties["login.google"] == "on" {
		fmt.Println("login google ativo")
	}

	shutdownFeature("login.google")
	fmt.Println(len(properties))

	delete(properties, "login.facebook")

	fmt.Println(len(properties))

	fmt.Println(isActive("login.google"))
}
