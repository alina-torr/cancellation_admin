DROP SCHEMA IF EXISTS hotelscheme CASCADE;
DROP SCHEMA IF EXISTS userscheme CASCADE;
DROP SCHEMA IF EXISTS bookingscheme CASCADE;


CREATE SCHEMA IF NOT EXISTS hotelscheme;
CREATE TABLE IF NOT EXISTS hotelscheme.hotel 
(
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text not null UNIQUE,
    country text not null,
    city text not null
    -- type text not null CHECK (type in ('City Hotel', 'Resort Hotel'))
);
-- CREATE TABLE IF NOT EXISTS hotelscheme.room 
-- (
--     id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
--     hotel_id int not null references hotelscheme.hotel(id),
--     price float4 not null
-- );

CREATE SCHEMA IF NOT EXISTS userscheme;
CREATE TABLE IF NOT EXISTS userscheme.manager 
(
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
    hotel_id int not null references hotelscheme.hotel(id),
    login text not null,
    password text not null
);
-- CREATE TABLE IF NOT EXISTS userscheme.client 
-- (
--     id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
--     phone_number text not null UNIQUE,
--     email text not null UNIQUE,
--     first_name text not null,
--     last_name text not null,
--     middle_name text,
--     country text
-- );




CREATE SCHEMA IF NOT EXISTS bookingscheme;
CREATE TABLE IF NOT EXISTS bookingscheme.booking 
(
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
    hotel_id int not null references hotelscheme.hotel(id),
    booking_id int not null,
    is_canceled boolean not null,
    lead_time int not null,
    arrival_date_year int not null,
    arrival_date_month text not null,
    arrival_date_week_number int not null,
    arrival_date_day_of_month int not null,
    stays_in_weekend_nights int not null,
    stays_in_week_nights int not null,
    adult_number int not null,
    child_number int not null,
    baby_number int not null,
    meal text CHECK (meal in ('BB', 'HB', 'FB', 'SC', 'Undefined')),
    country text,
    market_segment text CHECK (market_segment in ('Online TA', 'Offline TA/TO', 'Groups', 'Direct', 'Corporate', 'Aviation', 'Complementary')),       
    distribution_channel text CHECK (distribution_channel in ('TA/TO', 'Direct', 'Corporate', 'GDS')), 
    is_repeated_guest boolean not null,
    previous_cancellations int not null,
    previous_bookings_not_canceled int not null,
    -- reserved_room_id int not null references hotelscheme.room(id),
    reserved_room_type text not null,
    -- assigned_room_id int references hotelscheme.room(id),
    assigned_room_type text,
    booking_changes int not null,
    deposit_type text not null,
    agent text,
    company text,
    days_in_waiting_list int not null,
    customer_type text CHECK (customer_type in ('Transient', 'Transient-Party', 'Contract', 'Group')),
    adr int not null,
    required_car_parking_spaces int not null,
    total_of_special_requests int not null,
    -- special_requests text[],
    -- arrival_date date not null,
    -- booking_date date not null,
    -- confirmed_date date,
    -- night_stays int not null,
    
    -- deposit float4,
    reservation_status text CHECK (reservation_status in ('Created', 'Confirmed', 'Check-Out', 'Canceled', 'No-Show'))
);