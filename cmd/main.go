package main

func main() {
	server, err := InitHttpServer()
	if err != nil {
		panic(err)
	}
	server.Run()
}
