package main

func main() {
	r := setupRouter()

	// Listen and Serve on 0.0.0.0:8080
	r.Run(":8080")
}
