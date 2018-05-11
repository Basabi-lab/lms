[![CircleCI](https://circleci.com/gh/Basabi-lab/lms.svg?style=svg)](https://circleci.com/gh/Basabi-lab/lms)

PORT=8080 CLEARDB_DATABASE_URL="mysql://$(user):$(password)@/$(dbname)" ./lms

# For Docker
```
$ sudo docker pull mysql
$ sudo docker run -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -d mysql --default-authentication-plugin=mysql_native_password
$ sudo docker exec -it $(container_id) bash
% mysql -uroot
% > CREATE DATABASE lms;
% logout
$ PORT=8080 CLEARDB_DATABASE_URL="mysql://root:@/lms" ./lms &
$ npm run dev -- -p 3000 # Off CORS detection
```
