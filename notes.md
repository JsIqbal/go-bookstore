go mod is like the package.json in golang
gorm is a orm for backend in golang
-> go get "github.com/jinzhu/gorm"
then
-> go get "github.com/jinzhu/gorm/dialects/mysql"
-> go get "github.com/gorilla/mux"
//  routes > controllers > models

project flow : bookstore-routes,app.go, utils.go, book.go,  main.go , book.go, book-controller.go

db config for app.go file:
for go to connect:
mysql -u username -p
USE mysql;
SELECT user, host, authentication_string FROM user;
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'your_password';

CREATE DATABASE bookstore; // the db to connect from go app
GRANT ALL PRIVILEGES ON bookstore.* TO 'root'@'localhost';
