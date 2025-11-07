"""
SQLAlchemy ORM models representing database tables.
"""

from sqlalchemy import Column, Integer, String, Float, ForeignKey, CheckConstraint
from sqlalchemy.orm import relationship
from .database import Base


class Recipe(Base):
    """Represents a recipe entity."""
    __tablename__ = "recipes"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, nullable=False)
    instructions = Column(String, nullable=True)

    # One-to-many relationship: one recipe → many recipe items.
    items = relationship(
        "RecipeItem",
        back_populates="recipe",
        cascade="all, delete-orphan"
    )


class BaseIngredient(Base):
    """Represents a basic ingredient such as flour, water, salt, etc."""
    __tablename__ = "base_ingredients"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, nullable=False)
    unit = Column(String, nullable=True)  # Example: "g", "ml", "piece"


class RecipeItem(Base):
    """
    Represents a component of a recipe.
    Can reference either a base ingredient or another recipe.
    """
    __tablename__ = "recipe_items"

    id = Column(Integer, primary_key=True, index=True)
    recipe_id = Column(Integer, ForeignKey("recipes.id", ondelete="CASCADE"), nullable=False)
    item_type = Column(String, CheckConstraint("item_type IN ('base', 'recipe')"), nullable=False)
    item_id = Column(Integer, nullable=False)
    quantity = Column(Float, nullable=False)
    unit = Column(String, nullable=True)

    # Relationship back to parent recipe.
    recipe = relationship("Recipe", back_populates="items")
