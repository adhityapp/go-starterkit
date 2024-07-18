CREATE TABLE accounts (
  user_id SERIAL PRIMARY KEY, 
  username VARCHAR (50) UNIQUE NOT NULL, 
  password VARCHAR (50) NOT NULL, 
  email VARCHAR (255) UNIQUE NOT NULL, 
  created_at TIMESTAMP NOT NULL, 
  last_login TIMESTAMP
);

CREATE TABLE employee (
  firstname VARCHAR (50) NOT NULL, 
  lastname VARCHAR (50) NOT NULL, 
  employee_id SERIAL PRIMARY KEY, 
  hiredate TIMESTAMP NOT NULL, 
  terminationdate TIMESTAMP,
  salary VARCHAR (50) NOT NULL
);

CREATE TABLE annualreview (
  anid SERIAL PRIMARY KEY, 
  employee_id integer NOT NULL,
  reviewdate TIMESTAMP NOT NULL 
);


