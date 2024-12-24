-- create database sample_database;
use  sample_database;
desc countries;

create table regions(
regions_id int not null,
regions_name varchar(50) default null,
primary key(regions_id)
);

create table countries(
country_id int not null,
country_name varchar(50) default null,
regions_id int default null,
primary key(country_id),
foreign key(regions_id) references regions(regions_id)
);

create table locations(
location_id int not null,
street_address varchar(50) default null,
postal_code varchar(50) default null,
city varchar(50) default null,
state_province varchar(50) default null,
country_id int default null,
primary key(location_id),
foreign key(country_id) references countries(country_id)
);
show tables;

create table departments(
department_id int not null,
department_name varchar(50) default null,
location_id int default null,
primary key (department_id),
foreign key (location_id) references locations(location_id)
);


create table jobs(
job_id int not null,
job_title varchar(50) default null,
min_salary int default null,
max_salary int default null,
primary key(job_id) 
);

create table employees(
employee_id int not null,
first_name varchar(50) default null,
last_name varchar(50) default null,
email varchar(50) default null,
phone_number varchar(50) default null,
hire_date date default null,
job_id int default null,
salary int default null,
manager_id int default null,
department_id int default null,
primary key (employee_id),
foreign key (department_id) references departments (department_id),
foreign key (job_id)  references jobs(job_id),
foreign key (manager_id) references employees(employee_id)
);
-- desc employees;
create table dependents(
dependent_id int not null,
first_name varchar(50) default null,
last_name varchar(50) default null,
relationship varchar(50) default null,
employee_id int default null,
primary key(dependent_id),
foreign key(employee_id) references employees(employee_id)
);