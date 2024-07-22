-- grant_permissions.sql

-- Grant usage on schema public to local user
GRANT USAGE ON SCHEMA public TO local;

-- Grant create on schema public to local user
GRANT CREATE ON SCHEMA public TO local;

-- Grant all privileges on all tables in schema public to local user
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO local;

-- Grant all privileges on all sequences in schema public to local user
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO local;

-- Grant all privileges on the farrago database to local user
GRANT ALL PRIVILEGES ON DATABASE farrago TO local;

