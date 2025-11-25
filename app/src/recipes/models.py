"""
SQLAlchemy ORM model representing a simple recipe.
"""

from sqlalchemy import Column, Integer, String, Float
from sqlalchemy.orm import Session
from src.database import Base


class Recipe(Base):
    __tablename__ = "recipes"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, nullable=False)
    description = Column(String, nullable=True)
    quantity = Column(Integer, nullable=False)
    unit = Column(String, nullable=False)
    price = Column(Float, nullable=True)
    currency = Column(String, nullable=True)
    time = Column(Integer, nullable=True)
    difficulty = Column(String, nullable=False)
    img_url = Column(String, nullable=True)

    # -------- CLASSMETHODS -------- #

    @classmethod
    def create(cls, db: Session, data) -> "Recipe":
        """Create and persist a new recipe."""
        recipe = cls(
            name=data.name,
            description=data.description,
            quantity=data.quantity,
            unit=data.unit,
            price=data.price,
            currency=data.currency,
            time=data.time,
            difficulty=data.difficulty,
            img_url=data.img_url,
        )
        db.add(recipe)
        db.commit()
        db.refresh(recipe)
        return recipe

    @classmethod
    def fetch(cls, db: Session) -> list["Recipe"]:
        """Return all recipes."""
        return db.query(cls).all()

    @classmethod
    def get(cls, db: Session, recipe_id: int) -> "Recipe | None":
        """Return a recipe by ID."""
        return db.query(cls).filter(cls.id == recipe_id).first()

    @classmethod
    def delete(cls, db: Session, recipe_id: int) -> bool:
        """Delete a recipe by ID. Returns True if deleted, False otherwise."""
        recipe = cls.get(db, recipe_id)
        if not recipe:
            return False
        db.delete(recipe)
        db.commit()
        return True
