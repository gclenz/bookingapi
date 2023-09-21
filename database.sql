create table if not exists users (
  id uuid primary key default(gen_random_uuid()) not null,
  first_name varchar(25) not null,
  last_name varchar(25) not null,
  email varchar(45) unique not null,
  phone varchar(45) unique not null,
  document varchar(14) unique not null,
  date_of_birth date not null,
  role varchar(15) not null,
  password varchar(255) not null,
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