-- postgres
CREATE TABLE accounts (
  user_id SERIAL PRIMARY KEY, 
  username VARCHAR (50) UNIQUE NOT NULL, 
  password VARCHAR (50) NOT NULL, 
  email VARCHAR (255) UNIQUE NOT NULL, 
  created_at TIMESTAMP NOT NULL, 
  last_login TIMESTAMP
);

-- mysql
CREATE TABLE invoice (
    invoiceid SERIAL PRIMARY KEY,
    issuedate DATETIME NOT NULL,
    subject VARCHAR(50) NOT NULL,
    totalitems INT NOT NULL,
    customerid INT NOT NULL,
    amount FLOAT NOT NULL,
    duedate DATETIME NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'Unpaid'
);

CREATE TABLE customer (
    customerid SERIAL PRIMARY KEY,
    customername VARCHAR(100) NOT NULL,
    customeraddress VARCHAR(250) NOT NULL
);

CREATE TABLE items (
    itemid SERIAL PRIMARY KEY,
    itemname VARCHAR(50) NOT NULL,
    itemtype VARCHAR(50) NOT NULL,
    unitprice FLOAT NOT NULL
);

CREATE TABLE detail_items (
    detailitemid SERIAL PRIMARY KEY,
    invoiceid INT NOT NULL,
    itemid INT NOT NULL,
    qty INT,
    amount FLOAT
);


