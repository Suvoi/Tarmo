from sqlalchemy.orm import Session, selectinload
from .models import Recipe, Step


# --- Recipe CRUD ---

def get_recipe(db: Session, recipe_id: int) -> Recipe | None:
    """Fetch a single recipe by ID."""
    return db.query(Recipe).filter(Recipe.id == recipe_id).first()


def get_recipes(db: Session) -> list[Recipe]:
    """Fetch all recipes with steps eagerly loaded."""
    return db.query(Recipe).options(selectinload(Recipe.steps)).all()


def get_drafts(db: Session) -> list[Recipe]:
    """Fetch all draft recipes."""
    return db.query(Recipe).filter(Recipe.is_draft == True).all()


def update_recipe(db: Session, recipe: Recipe) -> Recipe:
    """Update an existing recipe."""
    db.commit()
    db.refresh(recipe)
    return recipe


def create_recipe(db: Session, recipe: Recipe) -> Recipe:
    """Insert a new recipe (does NOT commit - let caller handle transaction)."""
    db.add(recipe)
    db.flush()  # Generate ID but don't commit yet
    return recipe


def delete_recipe(db: Session, recipe: Recipe) -> None:
    """Delete a recipe (steps cascade automatically)."""
    db.delete(recipe)
    db.commit()


# --- Step CRUD ---

def get_step(db: Session, step_id: int) -> Step | None:
    """Fetch a single step by ID."""
    return db.query(Step).filter(Step.id == step_id).first()


def get_steps_by_recipe(db: Session, recipe_id: int) -> list[Step]:
    """Fetch all steps for a recipe, ordered."""
    return db.query(Step).filter(Step.recipe_id == recipe_id).order_by(Step.order).all()


def create_steps(db: Session, steps: list[Step]) -> None:
    """Bulk insert steps (does NOT commit - let caller handle transaction)."""
    db.add_all(steps)
    db.flush()


def delete_step(db: Session, step: Step) -> None:
    """Delete a single step."""
    db.delete(step)
    db.commit()