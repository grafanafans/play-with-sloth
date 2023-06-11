generate:
	docker run --rm --name sloth -v=$(CURDIR)/config:/sloth ghcr.io/slok/sloth  generate -i /sloth/slos/myservice.yaml -o /sloth/rules/myservice.yaml 
start:
	docker-compose up -d 
down:
	docker-compose down