-- create database DBCSuseraccount;
-- use DBCSuseraccount;
show databases;

create table DBCSuser(
userid int(11) unsigned auto_increment not null,
username varchar(20) not null,
passwd varchar(20) not null,
primary key(userid)
);

create table DBCSuser_detail(
userid int(11) unsigned auto_increment not null,
username varchar(20) not null,
useradder varchar(150) not null,
usermobile varchar(10) not null,
email varchar(50) not null,
primary key (userid)
);

show tables;
desc DBCSuser;
desc DBCSuser_detail;
drop table DBCSuser;
drop table DBCSuser_detail;


