run:
	go run main.go droptables
	go run main.go migrate
	go run main.go seed
	@echo "database restored sucessfully"