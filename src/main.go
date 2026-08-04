package main

func main()  {
	cfg := config{
		addr: ":8080",
	}
	api := api{
		config: cfg,
	}

	api.run(api.mount())
}
