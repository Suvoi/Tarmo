from sqlalchemy import ForeignKey, Column, Integer, String, Float
from sqlalchemy.orm import relationship, Session
from src.database import Base

class Step(Base):
    __tablename__ = "steps"

    id = Column(Integer, primary_key=True, index=True)
    recipe_id = Column(Integer, ForeignKey("recipes.id", ondelete="CASCADE"), nullable=False)
    order = Column(Integer, nullable=False)
    name = Column(String, nullable=False)
    instructions = Column(String, nullable=True)
    price = Column(Float, nullable=True)

    recipe = relationship("Recipe", back_populates="steps")

    # -------- CLASSMETHODS -------- #

    @classmethod
    def create(cls, db: Session, data) -> "Step":
        """Create and persist a new step."""
        step = cls(
            recipe_id=data.recipe_id,
            order=data.order,
            name=data.name,
            instructions=data.instructions,
            price=data.price,
        )
        db.add(step)
        db.commit()
        db.refresh(step)
        return step

    @classmethod
    def fetch(cls, db: Session, recipe_id: int) -> list["Step"]:
        """Return all steps for a given recipe, ordered by 'order'."""
        return db.query(cls).filter(cls.recipe_id == recipe_id).order_by(cls.order).all()

    @classmethod
    def get(cls, db: Session, step_id: int) -> "Step | None":
        """Return a step by ID."""
        return db.query(cls).filter(cls.id == step_id).first()

    @classmethod
    def update(cls, db: Session, step_id: int, **kwargs) -> "Step | None":
        """Update fields of a step by ID."""
        step = cls.get(db, step_id)
        if not step:
            return None
        for key, value in kwargs.items():
            if hasattr(step, key):
                setattr(step, key, value)
        db.commit()
        db.refresh(step)
        return step

    @classmethod
    def delete(cls, db: Session, step_id: int) -> bool:
        """Delete a step by ID. Returns True if deleted, False otherwise."""
        step = cls.get(db, step_id)
        if not step:
            return False
        db.delete(step)
        db.commit()
        return True
