"""
Entry point for the FastAPI application.
Defines routes and dependency injection for database sessions.
"""

from fastapi import FastAPI, Depends, HTTPException, status
from . import models, schemas, crud
from .database import Base, engine, SessionLocal

# Automatically create all tables on startup (safe for dev/prototyping)
Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="Tarmo",
    version="0.0.1",
    description="Optimizze and control batches with recipes."
)


# Dependency that provides a new database session per request.
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


@app.post("/recipes", response_model=schemas.RecipeResponse, status_code=status.HTTP_201_CREATED)
def create_recipe(data: schemas.RecipeCreate, db=Depends(get_db)):
    """Create a new recipe."""
    return crud.create_recipe(db, data)


@app.get("/recipes", response_model=list[schemas.RecipeResponse])
def list_recipes(db=Depends(get_db)):
    """Retrieve all recipes."""
    return crud.list_recipes(db)


@app.get("/recipes/{recipe_id}", response_model=schemas.RecipeResponse)
def get_recipe(recipe_id: int, db=Depends(get_db)):
    """Retrieve a recipe by its ID."""
    recipe = crud.get_recipe(db, recipe_id)
    if not recipe:
        raise HTTPException(status_code=404, detail="Recipe not found")
    return recipe


@app.post("/recipes/{recipe_id}/items", status_code=status.HTTP_201_CREATED)
def add_recipe_item(recipe_id: int, item: schemas.RecipeItemCreate, db=Depends(get_db)):
    """Attach an ingredient or subrecipe to a recipe."""
    # Optional: Validate recipe existence before adding items.
    if not crud.get_recipe(db, recipe_id):
        raise HTTPException(status_code=404, detail="Parent recipe not found")

    return crud.add_item_to_recipe(db, recipe_id, item)
