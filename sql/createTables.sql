CREATE TABLE Roles
(
    role_id   SERIAL PRIMARY KEY,
    role_name VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE Regions
(
    regions_id  SERIAL PRIMARY KEY,
    region_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE Cities
(
    city_id   SERIAL PRIMARY KEY,
    city_name VARCHAR(20) NOT NULL UNIQUE,
    region_id INT REFERENCES Regions (regions_id)
);

CREATE TABLE Users
(
    user_id      SERIAL PRIMARY KEY,
    first_name   VARCHAR(20)  NOT NULL,
    last_name    VARCHAR(40)  NOT NULL,
    fathers_name VARCHAR(40)  NOT NULL,
    email        VARCHAR(80)  NOT NULL UNIQUE,
    password     VARCHAR(256) NOT NULL,
    role_id      INT REFERENCES Roles (role_id),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active       BOOL      DEFAULT TRUE,
    updated_at   TIMESTAMP,
    deleted_at   TIMESTAMP
);

CREATE TABLE Cars
(
    car_id     SERIAL PRIMARY KEY,
    model      VARCHAR(50) NOT NULL,
    mark       VARCHAR(50) NOT NULL,
    autobody   VARCHAR(50) NOT NULL,
    car_number VARCHAR(10) NOT NULL,
    seats      VARCHAR(50) NOT NULL,
    user_id    INT REFERENCES Users (user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active     BOOL      DEFAULT TRUE,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE Statuses
(
    status_id   SERIAL PRIMARY KEY,
    status_name VARCHAR(20) NOT NULL
);

CREATE TABLE Seats
(
    seats_id     SERIAL PRIMARY KEY,
    car_id       INT REFERENCES Cars (car_id),
    seat_number  INT NOT NULL,
    is_available BOOL DEFAULT TRUE
);

CREATE TABLE Bookings
(
    bookings_id     SERIAL PRIMARY KEY,
    user_id         INT REFERENCES Users (user_id),
    seats_id        INT REFERENCES Seats (seats_id),
    status_id       INT REFERENCES Statuses (status_id),
    price           NUMERIC(10, 2),
    start_region_id INT REFERENCES Regions (regions_id),
    end_region_id   INT REFERENCES Regions (regions_id),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active          BOOL      DEFAULT TRUE,
    updated_at      TIMESTAMP,
    cancelled_at    TIMESTAMP
);

CREATE TABLE Chats
(
    chats_id SERIAL PRIMARY KEY,
    user_id  INT REFERENCES Users (user_id),
    mate_id  INT REFERENCES Users (user_id)
);

CREATE TABLE Messages
(
    message_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    from_user  INT REFERENCES Users (user_id),
    to_user    INT REFERENCES Users (user_id),
    message    TEXT NOT NULL,
    read       BOOL      DEFAULT FALSE,
    sent_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    read_at    TIMESTAMP,
    chat_id    INT REFERENCES Chats (chats_id)
);

CREATE TABLE Notifications_type
(
    type_id   SERIAL PRIMARY KEY,
    type_name VARCHAR(20) NOT NULL
);

-- Create Notifications table
CREATE TABLE Notifications
(
    notification_id   SERIAL PRIMARY KEY,
    user_id           INT REFERENCES Users (user_id),
    notification_type INT REFERENCES Notifications_type (type_id),
    description       TEXT NOT NULL
);

CREATE TABLE Routes
(
    route_id   INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    from_city  INT REFERENCES Cities (city_id),
    to_city    INT REFERENCES Cities (city_id),
    price      NUMERIC(10, 2),
    date       date,
    driver_id  INT REFERENCES Users (user_id),
    car_id     INT REFERENCES Cars (car_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active     BOOL      DEFAULT TRUE,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE DriverExtraInfo
(
    info_id   SERIAL PRIMARY KEY,
    driver_id INT REFERENCES Users (user_id),
    from_city INT REFERENCES Cities (city_id),
    to_city   INT references cities (city_id),
    quantity  INT
);

alter table users drop column fathers_name;
drop table roles;

create table personal_information
(
    info_id      INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id      int references users (user_id),
    first_name   varchar(20) not null,
    last_name    varchar(20) not null,
    fathers_name varchar(20) not null,
    about_me     text        not null,
    sex          varchar(10) not null,
    photo        text,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active       BOOL      DEFAULT TRUE,
    updated_at   TIMESTAMP,
    deleted_at   TIMESTAMP
);

alter table cities drop column region_id;
drop table regions;

insert into cities (city_name) values ('Dushanbe'), ('Kulob'), ('Khorog'), ('Khujand');
delete from seats where seats_id = 2

drop table Bookings;


CREATE TABLE Bookings
(
    bookings_id     SERIAL PRIMARY KEY,
    user_id         INT REFERENCES Users (user_id),
    driver_id       INT REFERENCES Users (user_id),
    seats_id        INT REFERENCES Seats (seats_id),
    status_id       INT REFERENCES Statuses (status_id),
    price           NUMERIC(10, 2),
    start_city_id INT REFERENCES Regions (regions_id),
    end_city_id   INT REFERENCES Regions (regions_id),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active          BOOL      DEFAULT TRUE,
    updated_at      TIMESTAMP,
    cancelled_at    TIMESTAMP
);