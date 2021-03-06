build_api_bin:
	go build -o test-task-api/bin/test-task-api test-task-api/main.go

run_api:
	cd test-task-api && go get -d ./... && go run main.go

run_ui:
	cd test-task-ui/src && npm install && npm start

run_server:
	make run_api & make run_ui