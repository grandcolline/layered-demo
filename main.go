package main

func main() {
	defer mysqlClose()
	serve()
}
