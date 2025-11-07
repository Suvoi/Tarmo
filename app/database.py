"""
Database configuration module.
Uses SQLite for local persistence, easily replaceable with PostgreSQL later.
"""

from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base

DATABASE_URL = "sqlite:///./data/recipes.db"

# SQLite requires `check_same_thread=False` for use with FastAPI's async execution model.
engine = create_engine(DATABASE_URL, connect_args={"check_same_thread": False})

# Session factory — creates independent database sessions per request.
SessionLocal = sessionmaker(bind=engine, autoflush=False, autocommit=False)

# Base class for declarative models.
Base = declarative_base()
