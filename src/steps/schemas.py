from pydantic import BaseModel
from typing import Optional

class StepBase(BaseModel):
    order: int
    name: str
    instructions: Optional[str] = None

class StepCreate(StepBase):
    recipe_id: int

class StepUpdate(BaseModel):
    order: Optional[int] = None
    name: Optional[str] = None
    instructions: Optional[str] = None
    
class StepOut(StepBase):
    id: int
    recipe_id: int

    class Config:
        orm_mode = True
