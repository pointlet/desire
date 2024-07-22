# Use the official PostgreSQL image as the base
FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=local
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=farrago

# Expose the PostgreSQL port
EXPOSE 5432

# Define a volume for persistent data
VOLUME ["/var/lib/postgresql/data"]

# The base image already contains an entrypoint script that initializes the database
# with the above environment variables, so no further commands are needed.