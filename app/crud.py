"""
CRUD utility functions that encapsulate all database interactions.
Keeps business logic separate from route definitions.
"""

from sqlalchemy.orm import Session
from . import models, schemas


def create_recipe(db: Session, data: schemas.RecipeCreate) -> models.Recipe:
    """Create and persist a new recipe."""
    recipe = models.Recipe(name=data.name, instructions=data.instructions)
    db.add(recipe)
    db.commit()
    db.refresh(recipe)
    return recipe


def list_recipes(db: Session) -> list[models.Recipe]:
    """Return all recipes."""
    return db.query(models.Recipe).all()


def get_recipe(db: Session, recipe_id: int) -> models.Recipe | None:
    """Return a recipe by ID, or None if not found."""
    return db.query(models.Recipe).filter(models.Recipe.id == recipe_id).first()


def add_item_to_recipe(
    db: Session, recipe_id: int, item: schemas.RecipeItemCreate
) -> models.RecipeItem:
    """Attach an ingredient or subrecipe to an existing recipe."""
    recipe_item = models.RecipeItem(recipe_id=recipe_id, **item.dict())
    db.add(recipe_item)
    db.commit()
    db.refresh(recipe_item)
    return recipe_item
