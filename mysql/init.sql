CREATE DATABASE IF NOT EXISTS mydatabase;

USE mydatabase;

CREATE TABLE IF NOT EXISTS countries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    population BIGINT NOT NULL,
    region VARCHAR(255) NOT NULL,
    capital VARCHAR(255) NOT NULL
);

INSERT INTO countries (name, population, region, capital) VALUES 
('United States', 331002651, 'North America', 'Washington, D.C.'),
('Canada', 37742154, 'North America', 'Ottawa'),
('Brazil', 212559417, 'South America', 'Brasília'),
('United Kingdom', 67886011, 'Europe', 'London'),
('Germany', 83783942, 'Europe', 'Berlin'),
('Australia', 25409300, 'Oceania', 'Canberra'),
('China', 1439323776, 'Asia', 'Beijing'),
('India', 1380004385, 'Asia', 'New Delhi'),
('South Africa', 59308690, 'Africa', 'Pretoria');