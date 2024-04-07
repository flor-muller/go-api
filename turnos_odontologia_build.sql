CREATE DATABASE IF NOT EXISTS  turnos_odontologia;

USE turnos_odontologia;

DROP TABLE IF EXISTS odontologos;

CREATE TABLE odontologos (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    apellido VARCHAR(255),
    nombre VARCHAR(255),
    matricula VARCHAR(255)
);