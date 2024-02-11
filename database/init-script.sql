CREATE TABLE IF NOT EXISTS client (
    clientID INT AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS professional (
    professionalID INT AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(255),
    address VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS administrator (
    administratorID INT AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS hairSalon (
    hairSalonID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    city VARCHAR(255),
    postal_code VARCHAR(10),
    professionalID INT,
    FOREIGN KEY (professionalID) REFERENCES professional(professionalID)
);

CREATE TABLE IF NOT EXISTS employee (
    employeeID INT AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    hairSalonID INT,
    FOREIGN KEY (hairSalonID) REFERENCES hairSalon(hairSalonID) 
);

CREATE TABLE IF NOT EXISTS service (
    serviceID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    price DECIMAL(10,2),
    duration INT,
    hairSalonID INT,
    FOREIGN KEY (hairSalonID) REFERENCES hairSalon(hairSalonID)
);

CREATE TABLE IF NOT EXISTS availability (
    availabilityID INT AUTO_INCREMENT PRIMARY KEY,
    employeeID INT,
    day_of_week ENUM('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'),
    start_time TIME,
    end_time TIME,
    FOREIGN KEY (employeeID) REFERENCES employee(employeeID)
);

CREATE TABLE IF NOT EXISTS reservation (
    reservationID INT AUTO_INCREMENT PRIMARY KEY,
    employeeID INT,
    clientID INT,
    serviceID INT,
    hairSalonID INT,
    date DATETIME,
    status ENUM('confirmed', 'canceled', 'expired') DEFAULT 'confirmed',
    FOREIGN KEY (clientID) REFERENCES client(clientID),
    FOREIGN KEY (serviceID) REFERENCES service(serviceID),
    FOREIGN KEY (hairSalonID) REFERENCES hairSalon(hairSalonID),
    FOREIGN KEY (employeeID) REFERENCES employee(employeeID)
);

CREATE TABLE IF NOT EXISTS request (
    requestID INT AUTO_INCREMENT PRIMARY KEY,
    professionalID INT,
    salon_name VARCHAR(255),
    address VARCHAR(255),
    city VARCHAR(255),
    postal_code VARCHAR(10),
    request_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    request_status ENUM('pending', 'accepted', 'rejected') DEFAULT 'pending',
    FOREIGN KEY (professionalID) REFERENCES professional(professionalID)
);

INSERT INTO administrator (firstname, lastname, email, password) VALUES ('admin', 'admin', 'admin@planity.com', 'admin');
INSERT INTO administrator (firstname, lastname, email, password) VALUES ('dylan', 'lgvn', 'dylan@planity.com', 'admin');
INSERT INTO professional (firstname, lastname, email, phone, address, password) VALUES ('John', 'Doe', 'pro@planity.com', '0689766371', '1 Rue des Monstres, Paris, 75016', 'password');
INSERT INTO professional (firstname, lastname, email, phone, address, password) VALUES ('dwgv', 'dgwd', 'proff@planity.com', '0689766371', '1 Rue des Monstres, Paris, 75016', 'password');
INSERT INTO client (firstname, lastname, email, password) VALUES ('Jane', 'Doe', 'client@planity.com', 'password');



INSERT INTO request (professionalID, salon_name, address, city, postal_code, request_date, request_status) VALUES (1, 'Salon coiffure 1', '1 Rue des Triple Monstres', 'Paris', '75017', '2024-02-05 14:30:00', 'pending');
INSERT INTO hairSalon (name, address, city, postal_code, professionalID)
VALUES ('Salon de coiffure XYZ', '123 Rue de la Beauté', 'Paris', '75001', 1);
INSERT INTO hairSalon (name, address, city, postal_code, professionalID)
VALUES ('Salon de coiffure ABC', '123 Rue de la Beauté', 'Paris', '75001', 2);
INSERT INTO employee (firstname, lastname, hairSalonID)
VALUES ('Dylan', 'Lang', 1);
INSERT INTO employee (firstname, lastname, hairSalonID)
VALUES ('Lancelot', 'LeCon', 2);
INSERT INTO availability (employeeID, day_of_week, start_time, end_time) VALUES
(1, 'Monday', '08:00:00', '16:00:00'),
(1, 'Tuesday', '08:00:00', '16:00:00'),
(1, 'Wednesday', '08:00:00', '16:00:00'),
(1, 'Thursday', '08:00:00', '16:00:00'),
(1, 'Friday', '08:00:00', '16:00:00'),
(1, 'Saturday', '10:00:00', '14:00:00'),
(1, 'Sunday', '10:00:00', '14:00:00');

INSERT INTO availability (employeeID, day_of_week, start_time, end_time) VALUES
(2, 'Monday', '09:00:00', '17:00:00'),
(2, 'Tuesday', '09:00:00', '17:00:00'),
(2, 'Wednesday', '09:00:00', '17:00:00'),
(2, 'Thursday', '09:00:00', '17:00:00'),
(2, 'Friday', '09:00:00', '17:00:00'),
(2, 'Saturday', '11:00:00', '15:00:00'),
(2, 'Sunday', '11:00:00', '15:00:00');

INSERT INTO service (name, description, price, duration, hairSalonID) 
VALUES ('Coupe de cheveux', 'Une coupe de cheveux standard pour hommes ou femmes', 30.00, 30, 1);

INSERT INTO service (name, description, price, duration, hairSalonID) 
VALUES ('Coupe de cheveux', 'Une coupe de cheveux standard pour hommes ou femmes', 30.00, 30, 2);


INSERT INTO request (professionalID, salon_name, address, city, postal_code, request_date, request_status) VALUES (2, 'Salon coiffure 2', '1 Rue des Double Monstres', 'Paris', '75012', '2024-02-05 14:30:00', 'pending');
INSERT INTO request (professionalID, salon_name, address, city, postal_code, request_date, request_status) VALUES (2, 'Salon coiffure 3', '1 Rue des quadruple Monstres', 'Paris', '75015', '2024-02-05 14:30:00', 'pending');