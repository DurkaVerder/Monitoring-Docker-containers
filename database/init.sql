-- Tables and index for PostgreSQL database
CREATE TABLE ping_result (
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(15) NOT NULL UNIQUE,
    ping_time INT,
    date_successful_ping TIMESTAMP
);