run_db:
	@echo " > Start Postgres (via docker-compose)"
	docker-compose up  --remove-orphans --build

run_unstable:
	@echo " > Start Unstable service"
	go build -o build/unstable_build github.com/s2ar/unstable/cmd/unstable &&  ./build/unstable_build -c config/config.yml server

run_check:
	@echo " > Start Check service"
	go build -o build/check_build github.com/s2ar/unstable/cmd/check &&  ./build/check_build

	