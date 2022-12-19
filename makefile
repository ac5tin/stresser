test:
	go test -race -timeout 30s -v $$(go list ./... | grep -v /tests/mocks/)

bench:
	go test -bench=./... -count 5 -cpu 1 -benchmem -v -run=^$$(go list ./... | grep -v /tests/mocks/)



sec:
	gosec -exclude-dir 'tests/mocks' ./...

scan:
	trivy fs .
	grype dir:./ --add-cpes-if-none


lint:
	golangci-lint run
