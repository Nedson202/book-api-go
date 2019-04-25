default:
	@echo "=============building Local API============="
	docker build -f Dockerfile -t main .

start: default
	@echo "=============starting api locally============="
	docker-compose up -d

logs:
	docker-compose logs -f

stop:
	docker-compose down

test:
	go test -v -cover ./...

clean: stop
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f

dev:
	@echo "=============starting api in development mode============="
	compileDaemon -command="./book-api"