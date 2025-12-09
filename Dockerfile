FROM python:3.12-slim AS build
ENV PYTHONUNBUFFERED=1
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt
COPY src/ /app/src/

FROM python:3.12-slim
WORKDIR /app
COPY --from=build /usr/local /usr/local
COPY --from=build /app /app
EXPOSE 9136
CMD ["uvicorn", "src.main:app", "--host", "0.0.0.0", "--port", "9136", "--log-level", "info"]
