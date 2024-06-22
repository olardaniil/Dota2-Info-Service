
run-tests:
	go test -cover ./...

create-tests-report:
	go test ./... -coverprofile=test_coverage.txt
	go tool  cover -html .\test_coverage.txt -o test_coverage_report.html

swag:
	swag init -g .\cmd\app\main.go