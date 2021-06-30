
-- user database
create database user_db;
create table user_tab (
  id bigint(20) not null auto_increment comment 'pk',
  name varchar(200) not null default '' comment 'name',
  status tinyint default 0 comment '0-enable，1-disable，default 0',
  password varchar(200) not null default '' comment 'password, salted+hashed', 
  email varchar(200) not null default '' comment 'email address',
  created_at bigint(20) not null default 0 comment 'create_at',
  updated_at bigint(20) not null default 0 comment 'update_at',
  db_update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  primary key (id),
  unique key idx_name(name)    
) engine=InnoDB  default charset=utf8mb4 comment 'user info'