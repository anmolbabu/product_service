create database product_service;
create schema product_service;
create table product_service.products(id VARCHAR(64) PRIMARY KEY NOT NULL, name VARCHAR(32), brand VARCHAR(32), quantity integer, price decimal, category VARCHAR(32));