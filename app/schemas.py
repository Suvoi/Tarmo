"""
Pydantic models for request validation and response serialization.
"""

from pydantic import BaseModel
from typing import Optional, List


class RecipeItemCreate(BaseModel):
    item_type: str
    item_id: int
    quantity: float
    unit: Optional[str] = None


class RecipeCreate(BaseModel):
    name: str
    instructions: Optional[str] = None


class RecipeResponse(BaseModel):
    id: int
    name: str
    instructions: Optional[str] = None

    class Config:
        orm_mode = True  # Enables automatic ORM object → Pydantic model conversion
