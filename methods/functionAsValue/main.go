package main

func reformat(message string, formatter func(string) string)string{
	message = formatter(message)// The reassignment is must because of immutability of string in GO
	message = formatter(message)
	message = formatter(message)
	return "TEXTIO: "+message
}