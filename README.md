[![CircleCI](https://circleci.com/gh/Basabi-lab/lms.svg?style=svg)](https://circleci.com/gh/Basabi-lab/lms)
[![Coverage Status](https://coveralls.io/repos/github/Basabi-lab/lms/badge.svg?branch=master)](https://coveralls.io/github/Basabi-lab/lms?branch=master)

# Run Service

```
$ mysql -u root
Mariadb[(none)]> CREATE DATABASE lms;
Mariadb[(none)]> quit
$ make run
```

# Deploy Command
tests/songに配置されている音楽データをDBにデータをPOSTするコマンド
***今は、実行するたびに、音楽データをDBから論理削除して、すべての情報をPOSTしており、かつ、共通のartistとかでも、別のartistとしてPOSTしている***

```
$ make run &
$ make deploy-run
```

# For Docker
```
$ sudo docker pull mysql
$ sudo docker run -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -d mysql --default-authentication-plugin=mysql_native_password
$ sudo docker exec -it $(container_id) bash
% mysql -uroot
% > CREATE DATABASE lms;
% logout
$ make run &
$ npm run dev -- -p 3000 # Off CORS detection
```


