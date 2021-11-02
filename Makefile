run:
	docker-compose up  --remove-orphans --build

run_unstable:
	@echo " > Start Unstable service"
	go build -o build/unstable_build github.com/s2ar/unstable/cmd/unstable &&  ./build/unstable_build -c config/config.yml server

run_check:
	@echo " > Start Check service"
	go build -o build/check_build github.com/s2ar/unstable/cmd/check &&  ./build/check_build

lint:
	@echo " > Start lint"
	@golangci-lint run

generate:
	mockgen -destination=./internal/services/objectservice/mocks.go -source=./internal/services/objectservice/repositories.go -package=objectservice
	# mockgen -destination=./internal/handlers/callbackhandler/mocks.go -source=./internal/handlers/callbackhandler/services.go -package=callbackhandler
	
	mockgen -destination=./internal/service/opendota/mocks.go -source=./internal/service/opendota/service.go -package=opendota
	mockgen -destination=./internal/application/mocks.go -source=./internal/application/app.go -package=application