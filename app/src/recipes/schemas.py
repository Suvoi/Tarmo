from pydantic import BaseModel, ConfigDict

class RecipeBase(BaseModel):
    name: str
    description: str | None = None
    quantity: float
    unit: str
    price: float | None = None
    currency: str | None = None
    time: int | None = None
    difficulty: str
    img_url: str | None = None


class RecipeCreate(RecipeBase):
    pass


class RecipeOut(RecipeBase):
    id: int

    model_config = ConfigDict(from_attributes=True)
