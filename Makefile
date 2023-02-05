build:
	docker build -t songjiayang/sloth-myservice:0.0.1 .
generate:
	docker run --rm --name sloth -v=$(CURDIR)/config:/sloth ghcr.io/slok/sloth  generate -i /sloth/slos/myservice.yaml -o /sloth/rules/myservice.yaml 
start:
	docker-compose up -d 
down:
	docker-compose down