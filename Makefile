test-all:
	go test -race -cover
build: test
	go build -race
open-coverage:
	go test -race -coverprofile=cover.out && \
	go tool cover -html=cover.out
docker-db-crean:
	rm -rf ./examples/master/data/*
	rm -rf ./examples/slave/data/*
docker-master-slave-db-run:
	cd ./examples && \
	docker-compose up
docker-master-slave-db-initialize:
	cd ./examples && \
	bash ./build.sh && \
	docker exec mysql_master sh -c "export MYSQL_PWD=111; mysql -u root mydb -e 'create table code(code int); insert into code values (100), (200)'" && \
	docker exec mysql_slave sh -c "export MYSQL_PWD=111; mysql -u root mydb -e 'select * from code \G'"	
docker-mysql-login-master:
	cd ./examples && \
	docker exec -it mysql_master mysql -u root -p111 mydb
docker-mysql-login-slave:
	cd ./examples && \
	docker exec -it mysql_slave mysql -u root -p111 mydb