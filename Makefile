build:
	@go build -o ./bin/nseTool .

run:build
	@./bin/nseTool

run-t:build
	@./bin/nseTool -t ACC

run-p:build
	@./bin/nseTool -p portfolio.txt

run-h:build
	@./bin/nseTool -h