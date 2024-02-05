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
    professionalID INT,
    hairSalonID INT,
    FOREIGN KEY (professionalID) REFERENCES professional(professionalID)
);

CREATE TABLE IF NOT EXISTS service (
    serviceID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    price DECIMAL(10,2),
    duration INT,
    professionalID INT,
    FOREIGN KEY (professionalID) REFERENCES professional(professionalID)
);

CREATE TABLE IF NOT EXISTS availability (
    availabilityID INT AUTO_INCREMENT PRIMARY KEY,
    employeeID INT,
    day_of_week ENUM('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'),
    start_time TIME,
    end_time TIME,
    FOREIGN KEY (employeeID) REFERENCES employee(employeeID)
);

CREATE TABLE IF NOT EXISTS timeSlot (
    timeSlotID INT AUTO_INCREMENT PRIMARY KEY,
    date DATE,
    start_time TIME,
    end_time TIME,
    professionalID INT,
    employeeID INT,
    availabilityID INT,
    FOREIGN KEY (professionalID) REFERENCES professional(professionalID),
    FOREIGN KEY (employeeID) REFERENCES employee(employeeID),
    FOREIGN KEY (availabilityID) REFERENCES availability(availabilityID)
);

CREATE TABLE IF NOT EXISTS reservation (
    reservationID INT AUTO_INCREMENT PRIMARY KEY,
    timeSlotID INT,
    clientID INT,
    serviceID INT,
    reservation_date DATETIME,
    reservation_status ENUM('pending', 'confirmed', 'cancelled'),
    FOREIGN KEY (timeSlotID) REFERENCES timeSlot(timeSlotID),
    FOREIGN KEY (clientID) REFERENCES client(clientID),
    FOREIGN KEY (serviceID) REFERENCES service(serviceID)
);

CREATE TABLE IF NOT EXISTS request (
    requestID INT AUTO_INCREMENT PRIMARY KEY,
    professionalID INT,
    salon_name VARCHAR(255),
    address VARCHAR(255),
    city VARCHAR(255),
    postal_code VARCHAR(10),
    request_date DATETIME,
    request_status ENUM('pending', 'accepted', 'rejected'),
    FOREIGN KEY (professionalID) REFERENCES professional(professionalID)
);

INSERT INTO administrator (firstname, lastname, email, password) VALUES ('admin', 'admin', 'admin@planity.com', 'admin');
INSERT INTO administrator (firstname, lastname, email, password) VALUES ('dylan', 'lgvn', 'dylan@planity.com', 'admin');
INSERT INTO professional (firstname, lastname, email, phone, address, password) VALUES ('John', 'Doe', 'pro@planity.com', '0689766371', '1 Rue des Monstres, Paris, 75016', 'password');
INSERT INTO professional (firstname, lastname, email, phone, address, password) VALUES ('dwgv', 'dgwd', 'proff@planity.com', '0689766371', '1 Rue des Monstres, Paris, 75016', 'password');
INSERT INTO client (firstname, lastname, email, password) VALUES ('Jane', 'Doe', 'client@planity.com', 'password');
INSERT INTO request (professionalID, salon_name, address, city, postal_code, request_date, request_status) VALUES (1, 'Salon coiffure 1', '1 Rue des Triple Monstres', 'Paris', '75017', '2024-02-05 14:30:00', 'pending');
INSERT INTO request (professionalID, salon_name, address, city, postal_code, request_date, request_status) VALUES (2, 'Salon coiffure 2', '1 Rue des Double Monstres', 'Paris', '75012', '2024-02-05 14:30:00', 'pending');
INSERT INTO request (professionalID, salon_name, address, city, postal_code, request_date, request_status) VALUES (2, 'Salon coiffure 3', '1 Rue des quadruple Monstres', 'Paris', '75015', '2024-02-05 14:30:00', 'pending');