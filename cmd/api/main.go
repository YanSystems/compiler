package main

func main() {
	api := Server{
		Port: "8000",
	}

	server := api.NewServer()
	api.Run(server)
}
