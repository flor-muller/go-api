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