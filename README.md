# flutter-golang-server
Its the server side of a project to demonstrate the full stack working of flutter as the front end and golang as the backend and postgresql as the database server

### SQL to create the database:
--------------------------------
CREATE DATABASE flutter-golang-server_dev01;

### SQL to create the tables:
-------------------------------- 
CREATE TABLE expenses (
  id SERIAL PRIMARY KEY,
  type VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  date DATE NOT NULL,
  rate NUMERIC
);