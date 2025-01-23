create table if not exists users (
  id uuid primary key default(gen_random_uuid()) not null,
  first_name varchar(25) not null,
  last_name varchar(25) not null,
  email varchar(45) unique not null,
  phone varchar(45) unique not null,
  document varchar(14) unique not null,
  date_of_birth date not null,
  role varchar(15) not null,
  code varchar(6) default('------') not null,
  code_expiration timestamp default(now()) not null,
  created_at timestamp default(now()) not null,
  updated_at timestamp default(now()) not null
);
-- drop table users

create table if not exists rooms (
  id uuid primary key default(gen_random_uuid()) not null,
  name varchar(50) not null,
  single_bed_count integer not null,
  double_bed_count integer not null,
  guests_limit integer not null,
  pet_friendly boolean not null,
  daily_price decimal(10,2) not null,
  created_at timestamp default(now()) not null,
  updated_at timestamp default(now()) not null
);
-- drop table rooms

create table if not exists bookings (
	id uuid primary key default(gen_random_uuid()) not null,
  room_id uuid references rooms(id) not null,
  customer_id uuid references users(id) not null,
  start_on timestamp not null,
  end_on timestamp not null,
  created_at timestamp default(now()) not null,
  updated_at timestamp default(now()) not null
);
-- drop table bookings

create table if not exists coupons (
	id uuid primary key default(gen_random_uuid()) not null,
  category varchar(10) not null,
  value decimal(10,2) not null,
  expires_in timestamp null,
  created_at timestamp default(now()) not null,
  updated_at timestamp default(now()) not null
);
-- drop table coupons

create table if not exists payments (
	id uuid primary key default(gen_random_uuid()) not null,
  external_id varchar(50) not null,
  booking_id uuid references bookings(id) not null,
  coupon_id uuid references coupons(id),
  gross_amount decimal(10,2) not null,
  net_amount decimal(10,2) not null,
  status varchar(50) not null,
  created_at timestamp default(now()) not null,
  updated_at timestamp default(now()) not null
);
-- drop table payments