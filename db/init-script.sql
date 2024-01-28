
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS Client (
    clientID INT AUTO_INCREMENT PRIMARY KEY,
    firstName VARCHAR(255),
    lastName VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Professional (
    professionalID INT AUTO_INCREMENT PRIMARY KEY,
    firstName VARCHAR(255),
    lastName VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(255),
    address VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Administrator (
    adminID INT AUTO_INCREMENT PRIMARY KEY,
    firstName VARCHAR(255),
    lastName VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS HairSalon (
    salonID INT AUTO_INCREMENT PRIMARY KEY,
    salon_name VARCHAR(255),
    address VARCHAR(255),
    city VARCHAR(255),
    postal_code VARCHAR(10),
    professionalID INT,
    FOREIGN KEY (professionalID) REFERENCES Professional(professionalID)
);

CREATE TABLE IF NOT EXISTS Employee (
    employeeID INT AUTO_INCREMENT PRIMARY KEY,
    firstName VARCHAR(255),
    lastName VARCHAR(255),
    professionalID INT,
    HairSalonID INT,
    FOREIGN KEY (professionalID) REFERENCES Professional(professionalID)
);

CREATE TABLE IF NOT EXISTS Service (
    serviceID INT AUTO_INCREMENT PRIMARY KEY,
    service_name VARCHAR(255),
    service_description TEXT,
    price DECIMAL(10,2),
    duration INT,
    professionalID INT,
    FOREIGN KEY (professionalID) REFERENCES Professional(professionalID)
);

CREATE TABLE IF NOT EXISTS Availability (
    availabilityID INT AUTO_INCREMENT PRIMARY KEY,
    employeeID INT,
    day_of_week ENUM('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'),
    start_time TIME,
    end_time TIME,
    FOREIGN KEY (employeeID) REFERENCES Employee(employeeID)
);

CREATE TABLE IF NOT EXISTS Time_Slot (
    timeSlotID INT AUTO_INCREMENT PRIMARY KEY,
    date DATE,
    start_time TIME,
    end_time TIME,
    professionalID INT,
    employeeID INT,
    availabilityID INT,
    FOREIGN KEY (professionalID) REFERENCES Professional(professionalID),
    FOREIGN KEY (employeeID) REFERENCES Employee(employeeID),
    FOREIGN KEY (availabilityID) REFERENCES Availability(availabilityID)
);

CREATE TABLE IF NOT EXISTS Reservation (
    reservationID INT AUTO_INCREMENT PRIMARY KEY,
    timeSlotID INT,
    clientID INT,
    serviceID INT,
    reservation_date DATETIME,
    reservation_status ENUM('pending', 'confirmed', 'cancelled'),
    FOREIGN KEY (timeSlotID) REFERENCES Time_Slot(timeSlotID),
    FOREIGN KEY (clientID) REFERENCES Client(clientID),
    FOREIGN KEY (serviceID) REFERENCES Service(serviceID)
);

CREATE TABLE IF NOT EXISTS Request (
    requestID INT AUTO_INCREMENT PRIMARY KEY,
    professionalID INT,
    salon_name VARCHAR(255),
    address VARCHAR(255),
    city VARCHAR(255),
    postal_code VARCHAR(10),
    request_date DATETIME,
    request_status ENUM('pending', 'accepted', 'rejected'),
    FOREIGN KEY (professionalID) REFERENCES Professional(professionalID)
);

GRANT ALL PRIVILEGES ON *.* TO 'user'@'%' WITH GRANT OPTION;