-- create database Sales;
-- use Sales;
create table customers(
customer_id varchar(50) not null,
customer_name varchar(50) default null,
city varchar(50) default  null,
order_count int default null,
primary key(customer_id)
);
-- insert into customers value ('02-1137231','Andie','Dampit',6),('15-0555193','Sutherlan','Dongshangguan',4),('17-9826606','Cybill','Huaibei',6),('23-0581020','Stephine','Bandera',9),('29-0813345','Jock','Frederiksberg',5),('31-7775741','Meredithe','Fengjiang',8),('36-4985773','Ginevra','Pajapita',1),('40-2510020','Stacey','Baisha',1),('41-5178437','Drucy','Trelleborg',3),('52-3535014','Izaak','Mugumu',5),('53-0525345','Reggie','Timóteo',2),('54-2364648','Lesya','Khorostkiv',4),('61-9760745','Frankie','Okuta',3),('66-1570441','Riobard','Clones',7),('71-4595008','Yettie','Mora',10),('71-7937313','Rodge','Stockholm',10),('72-1705612','Jillian','Luhansk',5),('72-9932424','Christie','Surulangun Rawas',6),('74-0511441','Osgood','Dobřany',8),('75-3153045','Lira','Tongliao',9),('77-7174697','Janka','Junyang',8),('84-5692031','Brenden','Sabanalarga',5),('90-7608988','Thor','As Suqaylibīyah',1),('94-7031101','Andie','Baitu',9),('95-8839983','Pippy','Sayama',2);


create table expired_products(
product_id int not null,
product_name varchar(50) default null,
expiration_date date default null,
primary key(product_id)
);
-- insert into expired_products value (1,'Sugar - Brown','2022-10-31'),(2,'Tea - Herbal I Love Lemon','2022-05-04'),(3,'Cloves - Ground','2023-03-26'),(4,'Trout - Rainbow, Frozen','2021-07-06'),(5,'Lychee','2021-12-14'),(6,'Yogurt - Strawberry, 175 Gr','2021-02-05'),(7,'Dates','2021-03-15'),(8,'Potatoes - Fingerling 4 Oz','2023-08-18'),(9,'Pork - Sausage, Medium','2023-06-08'),(10,'Chicken Giblets','2022-01-16');

create table products(
product_id int not null,
product_name varchar(50) default null,
unit_price float default null,
primary key(product_id)
);
-- insert into products value (1,'Sugar - Brown',66.52),(2,'Tea - Herbal I Love Lemon',10.73),(3,'Cloves - Ground',56.41),(4,'Trout - Rainbow, Frozen',54.27),(5,'Lychee',54.62),(6,'Yogurt - Strawberry, 175 Gr',81.97),(7,'Dates',32.79),(8,'Potatoes - Fingerling 4 Oz',35.62),(9,'Pork - Sausage, Medium',42.65),(10,'Chicken Giblets',12.21);

create table orders(
order_id int not null,
order_date date default null,
customer_id varchar(50) default null,
primary key (order_id),
foreign key(customer_id ) references customers(customer_id)
);
-- insert into orders value(1,'2022-01-17','61-9760745'),(2,'2021-09-07','52-3535014'),(3,'2022-02-16','02-1137231'),(4,'2023-07-04','84-5692031'),(5,'2021-07-25','54-2364648'),(6,'2022-02-22','95-8839983'),(7,'2022-07-30','40-2510020'),(8,'2021-11-26','36-4985773'),(9,'2023-11-25','90-7608988'),(10,'2023-03-05','71-4595008'),(11,'2021-06-04','71-7937313'),(12,'2023-02-14','77-7174697'),(13,'2021-08-01','15-0555193'),(14,'2021-05-14','31-7775741'),(15,'2023-01-23','41-5178437'),(16,'2021-05-28','17-9826606'),(17,'2023-04-10','72-9932424'),(18,'2021-05-26','72-1705612'),(19,'2021-04-04','75-3153045'),(20,'2021-12-18','94-7031101'),(21,'2021-05-23','53-0525345'),(22,'2023-04-07','66-1570441'),(23,'2021-12-03','23-0581020'),(24,'2023-02-13','29-0813345'),(25,'2021-11-30','74-0511441');

create table payments(
payment_id int not null,
payment_method varchar(50) default null,
amount int default null,
order_id int default null,
primary key(payment_id),
foreign key(order_id) references orders(order_id)
);

-- insert into payments value (1,'maestro',40,1),(2,'jcb',57,2),(3,'diners-club-enroute',73,3),(4,'jcb',30,4),(5,'china-unionpay',67,5),(6,'visa-electron',78,6),(7,'jcb',47,7),(8,'jcb',16,8),(9,'diners-club-carte-blanche',89,9),(10,'visa',32,10),(11,'jcb',16,11),(12,'diners-club-enroute',21,12),(13,'jcb',28,13),(14,'mastercard',51,14),(15,'visa',47,15),(16,'diners-club-enroute',44,16),(17,'jcb',24,17),(18,'mastercard',35,18),(19,'jcb',30,19),(20,'china-unionpay',85,20),(21,'visa-electron',37,21),(22,'solo',67,22),(23,'mastercard',45,23),(24,'jcb',18,24),(25,'jcb',68,25);

-- part 1
-- Q1
select * from customers where city = 'Mora';
-- Q2
-- insert into customers value ('07_1137077','Audi','SU-Land',0);
insert into customers value ('07-1137077','Audi','SU-Land',3);
-- select * from customers;
-- Q3
update customers set customer_name = 'Anna' where customer_id = '71-4595008';
-- set sql_safe_updates =1;
-- delete from customers where customer_id = '07_1137077';
-- Q4
select customer_id,max(order_date) as latest_your_order from orders group by  customer_id;
-- Q5
select customer_id,count(order_count) as total_order from customers group by  customer_id;
-- Q6
select product_name,unit_price from products where unit_price<20;
-- Q7
select city,avg(order_count) as avg_order_count from customers group by city;
-- Q8
select product_name,unit_price from products where unit_price >=20 and unit_price <=50;
-- Q9
select * from customers where customer_name = 'Anna';
-- Q10
select product_name from expired_products where expiration_date <'2023-01-01';
-- select * from expired_products;

-- part2
select * from customers;
-- Q1
alter table customers drop column order_count;
-- Q2
alter table customers change customer_name cus_name varchar(50);
-- Q3
select * from customers where cus_name not like 'A%' and cus_name not like 'B%'and cus_name not like 'C%';
-- Q4
select * from customers where customer_id like '%1%';
-- Q5
select customer_id,order_date from orders order by order_date desc;
-- Q6
select customer_id,order_date from orders order by order_date asc;
-- Q7
alter table orders modify column order_date varchar(60);
-- Q8
select count(product_id) as count from expired_products  where expiration_date <'2024-01-26';