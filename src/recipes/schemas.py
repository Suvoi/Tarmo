from pydantic import BaseModel, ConfigDict
from typing import List

from src.steps.schemas import StepBase
from src.steps.schemas import StepCreate

class RecipeBase(BaseModel):
    name: str
    description: str | None = None
    quantity: float
    unit: str
    difficulty: str
    img_url: str | None = None


class RecipeCreate(RecipeBase):
    steps: List[StepCreate]


class RecipeOut(RecipeBase):
    id: int
    steps: List[StepBase]
    model_config = ConfigDict(from_attributes=True)
