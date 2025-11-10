"""
Database configuration module.
Uses SQLite for local persistence, easily replaceable with PostgreSQL later.
"""

import os
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base

# Absolute path to the "app/data" folder
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))  # go up one level from src/
DATA_DIR = os.path.join(BASE_DIR, "data")
os.makedirs(DATA_DIR, exist_ok=True)  # ensure the folder exists

# Full path to the database
DATABASE_PATH = os.path.join(DATA_DIR, "recipes.db")
DATABASE_URL = f"sqlite:///{DATABASE_PATH}"

# Engine and session
engine = create_engine(DATABASE_URL, connect_args={"check_same_thread": False})
SessionLocal = sessionmaker(bind=engine, autoflush=False, autocommit=False)

# Base class for declarative models
Base = declarative_base()
