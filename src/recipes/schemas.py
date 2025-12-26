from pydantic import BaseModel, Field, ConfigDict
from typing import Optional


# -------- Step Schemas --------

class StepBase(BaseModel):
    name: str = Field(..., min_length=1, max_length=255)
    instructions: Optional[str] = None
    price: Optional[float] = Field(None, ge=0)


class StepCreate(StepBase):
    """Schema for creating a new step."""
    pass


class StepRead(StepBase):
    """Schema for reading a step (includes DB fields)."""
    id: int
    order: int
    recipe_id: int
    
    model_config = ConfigDict(from_attributes=True)


# -------- Recipe Schemas --------

class RecipeBase(BaseModel):
    name: str = Field(..., min_length=1, max_length=255)
    description: Optional[str] = None
    quantity: int = Field(..., gt=0)
    unit: str = Field(..., min_length=1, max_length=50)
    difficulty: str = Field(..., min_length=1, max_length=50)
    img_url: Optional[str] = Field(None, max_length=500)


class RecipeCreate(RecipeBase):
    """Schema for creating a new recipe with steps."""
    steps: list[StepCreate] = Field(..., min_length=1)


class RecipeRead(RecipeBase):
    """Schema for reading a recipe (includes DB fields and relationships)."""
    id: int
    is_draft: bool = False
    steps: list[StepRead] = []
    
    model_config = ConfigDict(from_attributes=True)


# -------- Draft Schemas --------

class RecipeDraftCreate(BaseModel):
    """Schema for creating a draft recipe (all fields optional except steps)."""
    name: Optional[str] = Field(None, max_length=255)
    description: Optional[str] = None
    quantity: Optional[int] = Field(None, gt=0)
    unit: Optional[str] = Field(None, max_length=50)
    difficulty: Optional[str] = Field(None, max_length=50)
    img_url: Optional[str] = Field(None, max_length=500)
    steps: list[StepCreate] = Field(default_factory=list)


class RecipeDraftRead(RecipeRead):
    """Schema for reading a draft recipe (inherits from RecipeRead)."""
    pass