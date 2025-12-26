from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session
from src.database import get_db
from . import repository as repo
from . import services
from .schemas import RecipeCreate, RecipeRead, RecipeDraftCreate, RecipeDraftRead


router = APIRouter()


# -------- Recipe Endpoints --------

@router.get("/", response_model=list[RecipeRead])
def read_recipes(db: Session = Depends(get_db)):
    """Return all recipes (calls repository directly - simple query)."""
    return repo.get_recipes(db)


@router.get("/{recipe_id}", response_model=RecipeRead)
def read_recipe(recipe_id: int, db: Session = Depends(get_db)):
    """Return a single recipe by ID (calls repository directly)."""
    recipe = repo.get_recipe(db, recipe_id)
    if not recipe:
        raise HTTPException(status_code=404, detail="Recipe not found")
    return recipe


@router.post(
    "/",
    response_model=RecipeRead,
    status_code=status.HTTP_201_CREATED,
)
def create_recipe(
    data: RecipeCreate,
    db: Session = Depends(get_db),
):
    """Create a new recipe with steps (calls service - complex operation)."""
    return services.create_recipe_with_steps(db, data)


@router.delete("/{recipe_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_recipe(recipe_id: int, db: Session = Depends(get_db)):
    """Delete a recipe by ID (calls service - includes validation)."""
    deleted = services.delete_recipe_by_id(db, recipe_id)
    if not deleted:
        raise HTTPException(status_code=404, detail="Recipe not found")


# -------- Draft Endpoints --------

@router.get("/drafts", response_model=list[RecipeDraftRead])
def read_drafts(db: Session = Depends(get_db)):
    """Return all draft recipes (calls repository directly - simple query)."""
    return repo.get_drafts(db)


@router.post(
    "/drafts",
    response_model=RecipeDraftRead,
    status_code=status.HTTP_201_CREATED,
)
def create_draft(
    data: RecipeDraftCreate,
    db: Session = Depends(get_db),
):
    """Create a new draft recipe (calls service - complex operation)."""
    return services.create_draft_recipe(db, data)