-- create database test
use test;
create table csuserlogin( 
userid int(9) unsigned auto_increment primary key not null,
username varchar(50) not null,
passwords varchar(50) not null
);

create table csuserlogin2( 
userid2 int(9) unsigned not null primary key,
username2 varchar(50) not null,
passwords2 varchar(50) not null);


create table csuserlogin3( 
userid2 int(9) not null ,
userid3 int(9) not null ,
username3 varchar(50) not null,
passwords3 varchar(50) not null,
primary key (userid3),
foreign key (userid2) references csuserlogin2(userid2));

-- drop table csuserlogin3;
show tables;
desc csuserlogin3;