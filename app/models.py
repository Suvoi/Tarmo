"""
SQLAlchemy ORM model representing a simple recipe.
"""

from sqlalchemy import Column, Integer, String
from .database import Base


class Recipe(Base):
    """Represents a simple recipe entity."""
    __tablename__ = "recipes"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, nullable=False)
    instructions = Column(String, nullable=True)  # Can store the full recipe text
