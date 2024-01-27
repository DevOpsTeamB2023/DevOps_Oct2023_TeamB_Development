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

CREATE TABLE IF NOT EXISTS `Record` (
`RecordID` int NOT NULL AUTO_INCREMENT,
`Name` varchar (50) NOT NULL,
`RoleOfContact` ENUM('Staff', 'Student'),
`NoOfStudents` int NOT NULL,
`AcadYr` varchar (10) NOT NULL,
`CapstoneTitle` varchar (50) NOT NULL,
`CompanyName` varchar (50) NOT NULL,
`CompanyContact` varchar (50) NOT NULL,
`ProjDesc` varchar (1000) NOT NULL,
PRIMARY KEY (`RecordID`)
)AUTO_INCREMENT=1;

INSERT INTO `Record` (`RecordID`, `Name`, `RoleOfContact`, `NoOfStudents`, `AcadYr`, `CapstoneTitle`, `CompanyName`, `CompanyContact`, `ProjDesc`)
VALUES(1, 'Zi Yi', 'Staff', 4, '2021/2022', 'Poverty Monitoring System', 'Shaniah Corporation', 'Koay YT', 'In the contemporary era, the intersection of virtual economies and real-world socio-economic issues has become increasingly relevant. This project aims to explore the correlation between spending patterns in the virtual world, specifically within the critically acclaimed MMORPG Final Fantasy XIV (With an expanded free trial which you can play through the entirety of A Realm Reborn and the award-winning Stormblood expansion up to level 70 for free with no restrictions on playtime?!!), and real-world poverty indicators (me).');

SELECT * FROM `Record`