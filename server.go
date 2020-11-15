package main 

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return "hello " + name
}

func main() {
	println(Hello("WWWWWWWWWWWWWW"))
	return
}
