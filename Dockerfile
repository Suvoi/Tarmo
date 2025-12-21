FROM python:3.12-slim AS build
WORKDIR /app
RUN apt-get update && apt-get install -y curl \
    && curl -sSL https://install.python-poetry.org | python3 -
ENV PATH="/root/.local/bin:$PATH"
COPY pyproject.toml poetry.lock* /app/
RUN poetry config virtualenvs.create false \
    && poetry install --no-dev --no-interaction --no-ansi
COPY src/ /app/src/

FROM python:3.12-slim
WORKDIR /app
COPY --from=build /usr/local /usr/local
COPY --from=build /app /app
EXPOSE 9136
CMD ["uvicorn", "src.main:app", "--host", "0.0.0.0", "--port", "9136", "--log-level", "info"]
