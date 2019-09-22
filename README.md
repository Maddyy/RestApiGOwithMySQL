# RestApiGOwithMySQL

sample of RestApi using Go with MySQL …
create table :-  CREATE TABLE `userinfo` (
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `city` varchar(100) DEFAULT NULL,
  `country` varchar(100) DEFAULT NULL,
  `phoneno` varchar(100) DEFAULT NULL,
  `salary` double DEFAULT NULL,
  PRIMARY KEY (`email`)
)
insert data:-
insert into `userinfo` values('sharad@gmail.com','test123','rahul','singh','mumbai','india',123,567);

http://localhost:8002/getUser/sharad@gmail.com&&test123
