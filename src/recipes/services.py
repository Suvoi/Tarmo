from sqlalchemy.orm import Session
from .models import Step, Recipe
from .repository import (
    get_step,
    get_steps_by_recipe,
    create_steps,
    delete_step,
    get_recipe,
    fetch_recipes,
    delete_recipe,
)

# -------- Queries --------

def get_step_by_id(db: Session, step_id: int) -> Step | None:
    """Return a single step by its ID."""
    return get_step(db, step_id)


def list_steps_for_recipe(db: Session, recipe_id: int) -> list[Step]:
    """Return all steps belonging to a specific recipe."""
    return get_steps_by_recipe(db, recipe_id)


def list_recipes(db: Session) -> list[Recipe]:
    """Return all recipes with their steps eagerly loaded."""
    return fetch_recipes(db)


def get_recipe_by_id(db: Session, recipe_id: int) -> Recipe | None:
    """Return a recipe by its ID."""
    return get_recipe(db, recipe_id)

# -------- Commands --------

def create_steps_for_recipe(
    db: Session,
    recipe_id: int,
    steps_data: list,
) -> list[Step]:
    """
    Create and persist steps for a given recipe.
    Step order is assigned automatically.
    """
    steps = [
        Step(
            recipe_id=recipe_id,
            order=i + 1,
            name=step.name,
            instructions=step.instructions,
            price=step.price,
        )
        for i, step in enumerate(steps_data)
    ]

    create_steps(db, steps)
    return steps


def delete_step_by_id(db: Session, step_id: int) -> bool:
    """Delete a step by ID. Returns True if deleted, False otherwise."""
    step = get_step(db, step_id)
    if not step:
        return False

    delete_step(db, step)
    return True


def create_recipe(db: Session, data) -> Recipe:
    """
    Create a recipe and delegate step creation to the step logic.
    """
    recipe = Recipe(
        name=data.name,
        description=data.description,
        quantity=data.quantity,
        unit=data.unit,
        difficulty=data.difficulty,
        img_url=data.img_url,
    )

    db.add(recipe)
    db.flush()  # obtain recipe.id before creating steps

    create_steps_for_recipe(
        db=db,
        recipe_id=recipe.id,
        steps_data=data.steps,
    )

    db.commit()
    db.refresh(recipe)
    return recipe


def delete_recipe_by_id(db: Session, recipe_id: int) -> bool:
    """
    Delete a recipe by ID.
    Related steps are removed automatically via cascade.
    """
    recipe = get_recipe(db, recipe_id)
    if not recipe:
        return False

    delete_recipe(db, recipe)
    return True
