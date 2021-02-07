build:
	docker build . -t aklinker1/url-shortener:dev
run: build
	docker-compose up
run-clean: build
	docker-compose up -V
deploy:
	heroku whoami &> /dev/null || heroku login
	heroku container:login
	docker build . -f Dockerfile.prod -t registry.heroku.com/apk-rip/web
	docker push registry.heroku.com/apk-rip/web
	heroku container:release web -a apk-rip