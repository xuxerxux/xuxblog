package xuxblog

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return "hello " + name
}

