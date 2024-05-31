DROP SCHEMA IF EXISTS hotelscheme CASCADE;
DROP SCHEMA IF EXISTS userscheme CASCADE;
DROP SCHEMA IF EXISTS bookingscheme CASCADE;

CREATE SCHEMA IF NOT EXISTS hotelscheme;
CREATE TABLE IF NOT EXISTS hotelscheme.hotel 
(
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text not null,
    country text not null,
    city text not null,
    api_key text not null
);

CREATE SCHEMA IF NOT EXISTS userscheme;
CREATE TABLE IF NOT EXISTS userscheme.manager 
(
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
    hotel_id int not null references hotelscheme.hotel(id),
    login text not null UNIQUE,
    password text not null
);

CREATE SCHEMA IF NOT EXISTS bookingscheme;
CREATE TABLE IF NOT EXISTS bookingscheme.booking 
(
    hotel_id int not null references hotelscheme.hotel(id),
    booking_id int not null,
    cancel_prediction float4 not null,
    arrival_date_year int not null,
    arrival_date_month text not null,
    arrival_date_day_of_month int not null,
    add_info text,
    primary key(hotel_id, booking_id)
);