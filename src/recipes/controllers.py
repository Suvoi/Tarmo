from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session

from src.database import get_db
from .services import (
    list_recipes,
    get_recipe_by_id,
    create_recipe,
    delete_recipe_by_id,
)
from .schemas import RecipeCreate, RecipeRead

router = APIRouter()


@router.get("/", response_model=list[RecipeRead])
def read_recipes(db: Session = Depends(get_db)):
    """Return all recipes."""
    return list_recipes(db)


@router.get("/{recipe_id}", response_model=RecipeRead)
def read_recipe(recipe_id: int, db: Session = Depends(get_db)):
    """Return a single recipe by ID."""
    recipe = get_recipe_by_id(db, recipe_id)
    if not recipe:
        raise HTTPException(status_code=404, detail="Recipe not found")
    return recipe


@router.post(
    "/",
    response_model=RecipeRead,
    status_code=status.HTTP_201_CREATED,
)
def create_new_recipe(
    data: RecipeCreate,
    db: Session = Depends(get_db),
):
    """Create a new recipe."""
    return create_recipe(db, data)


@router.delete("/{recipe_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_recipe(recipe_id: int, db: Session = Depends(get_db)):
    """Delete a recipe by ID."""
    deleted = delete_recipe_by_id(db, recipe_id)
    if not deleted:
        raise HTTPException(status_code=404, detail="Recipe not found")
