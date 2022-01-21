fmt:
	go fmt github.com/fuwn/bowl...

validate: .bowl/certificates/bowl.crt .bowl/certificates/bowl.key

run: fmt validate
	go run github.com/fuwn/bowl

build: fmt
	go build

ssl:
	mkdir -p .bowl/certificates
	openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes \
	  -out .bowl/certificates/bowl.crt \
	  -keyout .bowl/certificates/bowl.key \
	  -subj "/CN=bowl.fuwn.me"

docker: fmt
	docker build -t fuwn/bowl:latest .

dangling:
	sudo docker rmi $(sudo docker images -f "dangling=true" -q) --force
