// Package greet Use one word lower case for package name
// If possible save file as package name
package greet

// Hello function returns Hell:o + name
func Hello(name string, language string) string {
	if name == "" {
		name = "Dude"
	}
	return prefixGreeting(language) + name

}

func prefixGreeting(language string) (langHello string) {

	prefix := [3]string{
		"Hello ",
		"Hola ",
		"Bonjour ",
	}

	switch {
	case language == "espanhol":
		langHello = prefix[1]
	case language == "francês":
		langHello = prefix[2]
	default:
		langHello = prefix[0]
	}

	return

}
