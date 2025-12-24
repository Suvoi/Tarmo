from sqlalchemy import ForeignKey, Column, Integer, String, Float
from sqlalchemy.orm import relationship, Session
from src.database import Base

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

class Step(Base):
    __tablename__ = "steps"

    id = Column(Integer, primary_key=True, index=True)
    recipe_id = Column(Integer, ForeignKey("recipes.id", ondelete="CASCADE"), nullable=False)
    order = Column(Integer, nullable=False)
    name = Column(String, nullable=False)
    instructions = Column(String, nullable=True)
    price = Column(Float, nullable=True)

    recipe = relationship("Recipe", back_populates="steps")