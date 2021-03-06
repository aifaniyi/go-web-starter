run:
	docker-compose down -v
	docker-compose up -d postgres

	sleep 3
	export LOG_MODE=STDOUT && \
	cd cmd/server && \
		go build -o main && \
		./main