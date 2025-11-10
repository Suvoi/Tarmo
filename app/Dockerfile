# ---------- Base Stage ----------
FROM python:3.12-slim

WORKDIR /app

# Copy dependencies and install
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy source code
COPY app ./app
RUN mkdir -p /app/data

# Ensure an empty SQLite file exists (mounted later as a volume)
RUN touch /app/data/recipes.db

EXPOSE 9136

# Run FastAPI app with Uvicorn in production mode
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "9136"]
