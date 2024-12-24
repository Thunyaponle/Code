-- create database DBCSperson_infor;
-- use DBCSperson_infor;
-- show databases;

create table Csuserlogin(
userid int(9) unsigned auto_increment not null,
username varchar(50) not null,
password varchar(50) not null,
primary key(userid)
);


show tables;
desc Csuserlogin;

insert into Csuserlogin (username,password) value ('Jitdumrong Preechasuk','123456');
insert into Csuserlogin (username,password) value ('Ratchadaporn Kanawong','654321');
insert into Csuserlogin (username,password) value ('Data Science','01030507');
insert into Csuserlogin (username,password) value ('Computer Science','12131415');
insert into Csuserlogin (username,password) value ('information Engineering','34567890');


set sql_safe_updates=1;
DELETE FROM Csuserlogin;
DELETE FROM Csuserlogin WHERE username = 'kijf';
select username from Csuserlogin where password = '123456';
select * from Csuserlogin;
drop table Csuserlogin;

select * from Csuserlogin where userid = 2;
select userid,username from Csuserlogin where username = 'Computer Science' or username = 'information Engineering';
select password,username from Csuserlogin where username = 'Data Science' or username = 'information Engineering';
select userid,username from Csuserlogin where userid >=3 and userid <=5;
