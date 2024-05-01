-- Create the database
CREATE DATABASE product_catalog;

-- Switch to the newly created database
\c product_catalog;

-- Create the user
CREATE USER "product-catalog-user" WITH PASSWORD 'product-catalog-password';

-- Grant privileges to the user on the database
GRANT ALL PRIVILEGES ON DATABASE product_catalog TO "product-catalog-user";