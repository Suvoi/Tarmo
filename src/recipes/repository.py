from sqlalchemy.orm import Session, selectinload
from .models import Recipe, Step

# --- Recipe CRUD ---

def get_recipe(db: Session, recipe_id: int) -> Recipe | None:
    return db.query(Recipe).filter(Recipe.id == recipe_id).first()

def fetch_recipes(db: Session) -> list[Recipe]:
    recipes = db.query(Recipe).options(selectinload(Recipe.steps)).all()
    for recipe in recipes:
        recipe.steps.sort(key=lambda s: s.order)
    return recipes

def create_recipe(db: Session, recipe: Recipe) -> Recipe:
    db.add(recipe)
    db.commit()
    db.refresh(recipe)
    return recipe

def delete_recipe(db: Session, recipe: Recipe) -> None:
    db.delete(recipe)
    db.commit()


# --- Step CRUD ---

def get_step(db: Session, step_id: int) -> Step | None:
    return db.query(Step).filter(Step.id == step_id).first()

def get_steps_by_recipe(db: Session, recipe_id: int) -> list[Step]:
    steps = db.query(Step).filter(Step.recipe_id == recipe_id).order_by(Step.order).all()
    return steps

def create_steps(db: Session, steps: list[Step]) -> None:
    db.add_all(steps)
    db.commit()

def delete_step(db: Session, step: Step) -> None:
    db.delete(step)
    db.commit()
