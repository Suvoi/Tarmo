from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from src.database import get_db
from .models import Step
from .schemas import StepCreate, StepOut, StepUpdate

router = APIRouter(prefix="/steps", tags=["steps"])


@router.get("/recipe/{recipe_id}", response_model=list[StepOut])
def list_steps(recipe_id: int, db: Session = Depends(get_db)):
    """List all steps for a given recipe."""
    steps = Step.fetch(db, recipe_id)
    return steps


@router.get("/{step_id}", response_model=StepOut)
def get_step(step_id: int, db: Session = Depends(get_db)):
    """Get a step by ID."""
    step = Step.get(db, step_id)
    if not step:
        raise HTTPException(status_code=404, detail="Step not found")
    return step


@router.post("/", response_model=StepOut)
def create_step(data: StepCreate, db: Session = Depends(get_db)):
    """Create a new step."""
    return Step.create(db, data)


@router.put("/{step_id}", response_model=StepOut)
def update_step(step_id: int, data: StepUpdate, db: Session = Depends(get_db)):
    """Update an existing step."""
    step = Step.update(db, step_id, **data.dict(exclude_unset=True))
    if not step:
        raise HTTPException(status_code=404, detail="Step not found")
    return step


@router.delete("/{step_id}")
def delete_step(step_id: int, db: Session = Depends(get_db)):
    """Delete a step by ID."""
    ok = Step.delete(db, step_id)
    if not ok:
        raise HTTPException(status_code=404, detail="Step not found")
    return {"status": "deleted"}
