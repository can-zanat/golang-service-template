# golang-service-template
Basic golang api with MySql connection and unit tests

# this provide to create a docker container with mysql 8.0.23
docker run --name TestDB \
-e MYSQL_ROOT_USER=root \
-e MYSQL_ROOT_PASSWORD=12345 \
-e MYSQL_DATABASE=TestDB \
-p 3306:3306 \
-d mysql:8.0.23 \
--default-authentication-plugin=mysql_native_password \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_general_ci

# this provide to create a table in mysql
CREATE TABLE TestDB (
id INT NOT NULL AUTO_INCREMENT,
user_name VARCHAR(255) NOT NULL,
full_name VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL,
PRIMARY KEY (id)
);

# this provide to insert a data to mysql
INSERT INTO TestDB (user_name, full_name, email)
VALUES ('coniboy', 'Can Zanat', 'can.zanat@example.com');