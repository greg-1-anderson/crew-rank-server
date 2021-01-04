APP := crewrank-server

generate:
	go generate ./...

create-non-production-certs:
	mkdir -p certs
	echo '# Create CSR'
	openssl req  -new  -newkey rsa:2048  -nodes  -keyout certs/crewrank.key -out certs/crewrank.csr -batch -subj "/C=US/ST=California/L=San Francisco/O=CrewRank/OU=Service/CN=localhost"
	echo '# Create certificate'
	openssl  x509  -req  -days 365  -in certs/crewrank.csr  -signkey certs/crewrank.key  -out certs/crewrank.crt

create-production-certs:
	mkdir -p certs
	echo '# Create CSR'
	openssl req  -new  -newkey rsa:2048  -nodes  -keyout certs/crewrank.key  -out certs/crewrank.csr -batch -subj "/C=US/ST=California/L=San Francisco/O=CrewRank/OU=Service/CN=api.crew-rank.g1a.io"
	echo '# Create certificate'
	openssl  x509  -req  -days 365  -in certs/crewrank.csr  -signkey certs/crewrank.key  -out certs/crewrank.crt

show-certificate:
	openssl x509 -in certs/crewrank.crt -text

run:
	go run server.go

deploy:
	env GOOS=linux GOARCH=amd64 go build 
	scp -r certs crew-rank-server ga@radium.greenknowe.org:/home/ga
