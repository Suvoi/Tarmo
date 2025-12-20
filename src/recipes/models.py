"""
SQLAlchemy ORM model representing a simple recipe.
"""

from sqlalchemy import Column, Integer, String, Float
from sqlalchemy.orm import relationship, Session, selectinload
from src.database import Base

from src.steps.models import Step

class Recipe(Base):
    __tablename__ = "recipes"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, nullable=False)
    description = Column(String, nullable=True)
    quantity = Column(Integer, nullable=False)
    unit = Column(String, nullable=False)
    difficulty = Column(String, nullable=False)
    img_url = Column(String, nullable=True)

    steps = relationship(
        "Step",
        back_populates="recipe",
        cascade="all, delete-orphan",
        order_by="Step.order"
    )

    # -------- CLASSMETHODS -------- #

    @classmethod
    def create(cls, db: Session, data) -> "Recipe":
        """Create and persist a new recipe."""
        recipe = cls(
            name=data.name,
            description=data.description,
            quantity=data.quantity,
            unit=data.unit,
            difficulty=data.difficulty,
            img_url=data.img_url,
        )
        db.add(recipe)
        db.flush()

        steps = [
            Step(
                recipe_id=recipe.id,
                order=i+1,
                name=s.name,
                instructions=s.instructions
            )
            for i, s in enumerate(data.steps)
        ]

        db.add_all(steps)
        db.commit()
        db.refresh(recipe)
        return recipe

    @classmethod
    def fetch(cls, db: Session) -> list["Recipe"]:
        recipes = db.query(cls).options(selectinload(cls.steps)).all()
        for recipe in recipes:
            recipe.steps.sort(key=lambda s: s.order)
        return recipes

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
