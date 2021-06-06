build:
	docker build . -t aklinker1/url-shortener:dev
prod:
	./scripts/build.sh
run: build
	docker-compose up --remove-orphans
run-clean: build
	docker-compose up --remove-orphans -V

deploy:
	heroku whoami &> /dev/null || heroku login
	heroku container:login
	docker build . -t aklinker1/url-shortener:prod
	docker tag aklinker1/url-shortener:prod registry.heroku.com/apk-rip/web
	docker push registry.heroku.com/apk-rip/web
	heroku container:release -a apk-rip web
