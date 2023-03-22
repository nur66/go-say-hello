package go_say_hello

func SayHello(name string) string {
	return "Hello " + name
}

func SayGreeting(country string) string {
	if country == "Indonesia" {
		return "Selamat Pagi"
	} else if country == "Japan" {
		return "Konichiwa"
	} else if country == "France" {
		return "Bonjour"
	} else if country == "Saudi Arabia" {
		return "Sobahun Nur"
	} else {
		return "Good Morning"
	}
}
