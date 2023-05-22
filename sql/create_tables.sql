CREATE TABLE  roles (
    id SERIAL PRIMARY KEY NOT NULL,
    role_name varchar NOT NULL  
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    name varchar  NOT NULL,
    surname varchar NOT NULL,
    username varchar NOT NULL,
    email varchar NOT NULL,
    role_id INTEGER REFERENCES roles(id)
    
);
