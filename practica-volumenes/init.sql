CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL
);
INSERT INTO usuarios (nombre) VALUES ('Administrador Sistema');