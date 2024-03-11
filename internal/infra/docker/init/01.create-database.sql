-- Check if user 'ash' exists, if not, create with superuser privileges
DO
$do$
    BEGIN
        IF NOT EXISTS (
            SELECT FROM pg_catalog.pg_roles WHERE rolname = 'ash'
        ) THEN
            CREATE USER ash WITH PASSWORD 'strong_password' SUPERUSER;
        END IF;
    END
$do$;

-- Create database named 'dbname'
CREATE DATABASE dbname;

-- Connect to the new database and enable UUID extension.
\c dbname

-- Create the UUID extension.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
