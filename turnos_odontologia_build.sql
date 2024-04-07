-- Delete user if already exists
DROP USER IF EXISTS 'root'@'localhost';

-- Create user with all privileges
CREATE USER 'root'@'localhost' IDENTIFIED BY 'root';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost';

-- Create database
CREATE DATABASE IF NOT EXISTS  turnos_odontologia;

USE turnos_odontologia;

-- Create table odontologos
DROP TABLE IF EXISTS odontologos;

CREATE TABLE odontologos (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    apellido VARCHAR(255),
    nombre VARCHAR(255),
    matricula VARCHAR(255)
);

-- Create table pacientes
DROP TABLE IF EXISTS pacientes;

CREATE TABLE pacientes (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    apellido VARCHAR(255),
    nombre VARCHAR(255),
    domicilio VARCHAR(255),
    dni VARCHAR(255),
    alta VARCHAR(255)
);

-- Create table turnos
DROP TABLE IF EXISTS turnos;

CREATE TABLE turnos (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    id_paciente INTEGER,
    id_odontologo INTEGER,
    fecha VARCHAR(255),
    hora VARCHAR(255),
    descripcion VARCHAR(255),
    FOREIGN KEY (id_paciente) REFERENCES pacientes(id) ON DELETE CASCADE,
    FOREIGN KEY (id_odontologo) REFERENCES odontologos(id) ON DELETE CASCADE
);