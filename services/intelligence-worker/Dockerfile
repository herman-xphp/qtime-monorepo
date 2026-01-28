# Base Image
FROM python:3.14-slim
WORKDIR /app

# Install dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy App
COPY . .

# Run
CMD [ "uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]