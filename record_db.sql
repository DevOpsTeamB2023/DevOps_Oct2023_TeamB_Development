CREATE USER 'record_system'@'localhost' IDENTIFIED BY 'dopasgpwd';
GRANT ALL ON *.* TO 'record_system'@'localhost';

CREATE DATABASE IF NOT EXISTS `record_db` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `record_db`;

CREATE TABLE IF NOT EXISTS `Account` (
`AccID` int NOT NULL AUTO_INCREMENT,
`Username` varchar (50) NOT NULL,
`Password` varchar (50) NOT NULL,
`AccType` varchar (10) NOT NULL,
`AccStatus` varchar (30) NOT NULL,
PRIMARY KEY (`AccID`)
) ENGINE=InnoDB AUTO_INCREMENT=2002 DEFAULT CHARSET=utf8mb4;

INSERT INTO `Account` (`AccID`, `Username`, `Password`, `AccType`, `AccStatus`)
VALUES(1001, 'Shaniah', 'adminpwd1', 'Admin', 'Created'),
(2001, 'ziyi', 'userpwd1', 'User', 'Created'),
(2999, 'testdelete', 'deletetestpwd', 'User', 'Created'),
(2998, 'testapprove', 'approvetestpwd', 'User', 'Pending'),
(2997, 'testupdate', 'updatetestpwd', 'User', 'Created');

SELECT * FROM `Account`;