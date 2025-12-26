from sqlalchemy.orm import Session
from .models import Step, Recipe
from . import repository as repo


# -------- Recipe Commands (Business Logic) --------

def create_recipe_with_steps(db: Session, data) -> Recipe:
    """
    Create a complete recipe with its steps.
    Orchestrates recipe creation and step assignment.
    """
    recipe = Recipe(
        name=data.name,
        description=data.description,
        quantity=data.quantity,
        unit=data.unit,
        difficulty=data.difficulty,
        img_url=data.img_url,
        is_draft=False
    )
    repo.create_recipe(db, recipe)  # Flush to get ID
    
    if data.steps:
        _create_steps_for_recipe(db, recipe.id, data.steps)
    
    db.commit()
    db.refresh(recipe)
    return recipe


def create_draft_recipe(db: Session, data) -> Recipe:
    """
    Create a draft recipe (allows incomplete data).
    Business rule: drafts can have empty fields.
    """
    recipe = Recipe(
        name=data.name or "",
        description=data.description,
        quantity=data.quantity or 0,
        unit=data.unit or "",
        difficulty=data.difficulty or "",
        img_url=data.img_url,
        is_draft=True
    )
    repo.create_recipe(db, recipe)  # Flush to get ID
    
    if data.steps:
        _create_steps_for_recipe(db, recipe.id, data.steps)
    
    db.commit()
    db.refresh(recipe)
    return recipe


def delete_recipe_by_id(db: Session, recipe_id: int) -> bool:
    """
    Delete a recipe by ID with validation.
    Business rule: verify recipe exists before deletion.
    """
    recipe = repo.get_recipe(db, recipe_id)
    if not recipe:
        return False
    
    repo.delete_recipe(db, recipe)
    return True


# -------- Step Commands (Business Logic) --------

def delete_step_by_id(db: Session, step_id: int) -> bool:
    """
    Delete a step by ID with validation.
    Business rule: verify step exists before deletion.
    """
    step = repo.get_step(db, step_id)
    if not step:
        return False
    
    repo.delete_step(db, step)
    return True


# -------- Private Helpers --------

def _create_steps_for_recipe(db: Session, recipe_id: int, steps_data: list) -> list[Step]:
    """
    Internal helper: create steps with auto-assigned order.
    Business rule: step order is 1-indexed and sequential.
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
    repo.create_steps(db, steps)
    return steps