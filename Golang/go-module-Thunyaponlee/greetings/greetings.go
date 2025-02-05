package greetings

func Hello(name string) string {
	str := name
	if str == "" {
		return "Hello, World"
	} else {
		return "Hello, " + name
	}
}

/*func formatMessage(name string) string {
	str2 := name
	if str2 == "" {
		return "*** World ***"
	} else {
		return "*** " + name + "  ***"
	}
}*/

func formatMessage(name string) string {
	str2 := name
	if str2 == "" {
		name = "World"
	}
	return "*** " + name + " ***"

}

func FormatHello(name string) string {
	return "Hello, " + formatMessage(name)
}
