from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from ..database import get_db
from .models import Recipe
from .schemas import RecipeCreate, RecipeOut

router = APIRouter()


@router.get("/", response_model=list[RecipeOut])
def list_recipes(db: Session = Depends(get_db)):
    return Recipe.fetch(db)


@router.get("/{recipe_id}", response_model=RecipeOut)
def get_recipe(recipe_id: int, db: Session = Depends(get_db)):
    recipe = Recipe.get(db, recipe_id)
    if not recipe:
        raise HTTPException(status_code=404, detail="Recipe not found")
    return recipe


@router.post("/", response_model=RecipeOut)
def create_recipe(data: RecipeCreate, db: Session = Depends(get_db)):
    return Recipe.create(db, data)


@router.delete("/{recipe_id}")
def delete_recipe(recipe_id: int, db: Session = Depends(get_db)):
    ok = Recipe.delete(db, recipe_id)
    if not ok:
        raise HTTPException(status_code=404, detail="Recipe not found")
    return {"status": "deleted"}
