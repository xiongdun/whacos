FROM golang:1.12.6

# Create the directory where the application will reside
RUN mkdir /app

# Copy the application files (needed for production)
ADD whacos /app/whacos

# Set the working directory to the app directory
WORKDIR /app

# Expose the application on port 8080.
# This should be the same as in the app.conf file
EXPOSE 8090

# Set the entry point of the container to the application executable
ENTRYPOINT /app/whacos