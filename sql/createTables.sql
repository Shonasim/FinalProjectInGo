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
    first_name   VARCHAR(20) NOT NULL,
    last_name    VARCHAR(40) NOT NULL,
    fathers_name VARCHAR(40) NOT NULL,
    email        VARCHAR(80) NOT NULL UNIQUE,
    password     VARCHAR(256) NOT NULL,
    role_id      INT REFERENCES Roles (role_id),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active       BOOL DEFAULT TRUE,
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
    region_id  INT REFERENCES Regions (regions_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active     BOOL DEFAULT TRUE,
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
    active          BOOL DEFAULT TRUE,
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
    read       BOOL DEFAULT FALSE,
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

-- Create Routes table
CREATE TABLE Routes
(
    route_id   SERIAL PRIMARY KEY,
    from_city  INT REFERENCES Cities (city_id),
    to_city    INT REFERENCES Cities (city_id),
    price      NUMERIC(10, 2),
    date       DATE,
    user_id    INT REFERENCES Users (user_id),
    car_id     INT REFERENCES Cars (car_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active     BOOL DEFAULT TRUE,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE DriverExtraInfo
(
    info_id   SERIAL PRIMARY KEY,
    driver_id INT REFERENCES Users (user_id),
    from_city INT REFERENCES Cities (city_id),
    quantity  INT
);

INSERT INTO Roles (role_name) VALUES
  ('Driver'),
  ('Passenger');


INSERT INTO Regions (region_name) VALUES
  ('ГБАО'),
  ('Хатлонская область'),
  ('Согдийская область'),
  ('Города республиканского подчинения'),



INSERT INTO Cities (city_name, region_id) VALUES
  ('Хорог', 1),
  ('Куляб', 2),
  ('Худжанд', 3),
  ('Вахдат', 4),



INSERT INTO Users (first_name, last_name, fathers_name, email, password, role_id) VALUES
  ('Raimdodov', 'Shonasim', 'Umedovich', 'raimdodov.sh@gmail.com', '2003', 1),
  ('Kurush', 'Qosimi', 'Qosimovich', 'qosimikurush@gmail.com', '2123', 2);


INSERT INTO Cars (model, mark, autobody, car_number, seats, user_id, region_id) VALUES
    ('Tayota', 'Land Cruiser 100', 'Picap', '1234AA04', '5', 1, 1),
    ('Tayota', 'Land Cruiser 200', 'Picap', '1224AA01', '5', 1, 1);

INSERT INTO Statuses (status_name) VALUES
   ('Pending'),
   ('Confirmed'),
   ('Cancelled');

INSERT INTO Seats (car_id, seat_number, is_available) VALUES
  (1, 6, TRUE),
  (2, 6, TRUE);



INSERT INTO Bookings (user_id, seats_id, status_id, price, start_region_id, end_region_id) VALUES
   (2, 1, 1, 100.00, 1, 2);



INSERT INTO Chats (user_id, mate_id) VALUES
     (1, 2);



INSERT INTO Messages (from_user, to_user, message, chat_id) VALUES
    (2, 1, 'Привет, как дела?', 1);


INSERT INTO Notifications_type (type_name) VALUES
   ('Message'),
   ('Booking'),
   ('Alert');


INSERT INTO Notifications (user_id, notification_type, description) VALUES
    (2, 2, 'Ваше бронирование подтверждено.');


INSERT INTO Routes (from_city, to_city, price, date, user_id, car_id) VALUES
      (1, 2, 200.00, '2024-10-15', 2, 1);


INSERT INTO DriverExtraInfo (driver_id, from_city, quantity) VALUES
     (1, 1, 10);
