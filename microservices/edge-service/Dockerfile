# Use an official Python runtime as the base image
FROM python:3.7-slim-buster

# Set the working directory to /app
WORKDIR /app

# Copy the requirements.txt file to the container
COPY requirements.txt .

# Install the required packages
RUN pip install --no-cache-dir -r requirements.txt

# Copy the edge-service.py file to the container
COPY edge-service.py .

# Define the environment variable for Flask
ENV FLASK_APP=edge-service.py

# Run the command to start the Flask application
CMD ["flask", "run", "--host=0.0.0.0"]
