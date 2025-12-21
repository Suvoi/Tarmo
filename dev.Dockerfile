FROM ghcr.io/astral-sh/uv:python3.12-bookworm-slim
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV UV_VENV_DISABLE=1
WORKDIR /app
COPY pyproject.toml uv.lock /app/
RUN uv sync
EXPOSE 9136
CMD ["uv", "run", "uvicorn", "src.main:app", "--host", "0.0.0.0", "--port", "9136", "--reload","--log-level", "debug"]