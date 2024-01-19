# golang-service-template
Basic golang api with MySql connection and unit tests

docker run --name TestDB \
-e MYSQL_ROOT_USER=root \
-e MYSQL_ROOT_PASSWORD=12345 \
-e MYSQL_DATABASE=TestDB \
-p 3306:3306 \
-d mysql:8.0.23 \
--default-authentication-plugin=mysql_native_password \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_general_ci