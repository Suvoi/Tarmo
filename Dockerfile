FROM python:3.12-slim as build
ENV PYTHONUNBUFFERED=1
ENV POETRY_VIRTUALENVS_CREATE=false  # si usas poetry
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt
COPY src/ /app/

FROM python:3.12-slim
WORKDIR /app
COPY --from=build /usr/local/lib/python3.12/site-packages /usr/local/lib/python3.12/site-packages
COPY --from=build /app /app
EXPOSE 9136
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "9136", "--log-level", "info"]
