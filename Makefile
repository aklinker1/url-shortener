build:
	docker build . -t aklinker1/url-shortener:dev
run: build
	docker-compose up
run-clean: build
	docker-compose up -V
