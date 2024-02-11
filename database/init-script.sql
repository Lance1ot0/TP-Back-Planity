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
    interval_time INT,
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