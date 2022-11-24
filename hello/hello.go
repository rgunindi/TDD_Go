package main

const prefix = "hello"
var language = make(map[string]string)

func init() {
	language["Spanish"] = "Hola"
	language["French"] = "Bonjour"
}

var found = func(key string) bool { _, ok := language[key]; return ok }
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	switch {
	case found(lang):
		return language[lang] + ", " + name
	default:
		return prefix + " " + name
	}
}