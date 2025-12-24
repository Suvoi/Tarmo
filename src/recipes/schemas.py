from pydantic import BaseModel, Field
from typing import Optional, List

# -------- Step Schemas --------

class StepBase(BaseModel):
    name: str = Field(..., min_length=1)
    instructions: Optional[str] = None
    price: Optional[float] = None


class StepCreate(StepBase):
    pass


class StepRead(StepBase):
    id: int
    order: int
    recipe_id: int

    class Config:
        from_attributes = True


# -------- Recipe Schemas --------

class RecipeBase(BaseModel):
    name: str = Field(..., min_length=1)
    description: Optional[str] = None
    quantity: int = Field(..., gt=0)
    unit: str
    difficulty: str
    img_url: Optional[str] = None


class RecipeCreate(RecipeBase):
    steps: List[StepCreate]


class RecipeRead(RecipeBase):
    id: int
    steps: List[StepRead]

    class Config:
        from_attributes = True

