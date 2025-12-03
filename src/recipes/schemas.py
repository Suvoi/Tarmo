from pydantic import BaseModel, ConfigDict

class RecipeBase(BaseModel):
    name: str
    description: str | None = None
    instructions: str
    quantity: float
    unit: str
    difficulty: str
    img_url: str | None = None


class RecipeCreate(RecipeBase):
    pass


class RecipeOut(RecipeBase):
    id: int

    model_config = ConfigDict(from_attributes=True)
