# Docker

```shell
# Run database server
docker run -d --name mysql-200-lab --privileged=true -e MYSQL_ROOT_PASSWORD="1234" -e MYSQL_USER="food_delivery" -e MYSQL_PASSWORD="1234" -e MYSQL_DATABASE="food_delivery" -dp 3306:3306 bitnami/mysql:5.7
```
