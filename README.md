# gogin-practice
gin golang framework

# MYSQL DDL status employee
create table t_status(
statusid varchar(255) not null,
status varchar(255),
primary key(statusid)
);

# MYSQL data status employee
insert into t_status values('1','Employee');
insert into t_status values('2','Contract');
insert into t_status values('3','Freelance');

# MYSQL DDL user
CREATE TABLE t_user ( 
	userid varchar(255) NOT NULL, 
    username varchar(255) NOT NULL, 
    name varchar(255) DEFAULT NULL, 
    statusid varchar(255) not null,
PRIMARY KEY (userid) ,
foreign key(statusid) references t_status(statusid)
)
