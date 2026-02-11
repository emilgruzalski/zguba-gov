"""Pydantic schemas for found items."""

from pydantic import BaseModel, EmailStr, Field


class MunicipalityInfo(BaseModel):
    """Municipality information."""

    name: str
    type: str
    contactEmail: EmailStr


class ItemInfo(BaseModel):
    """Found item details."""

    name: str
    category: str
    date: str
    location: str
    status: str = "available"
    description: str | None = None


class PickupInfo(BaseModel):
    """Item pickup conditions."""

    deadline: int
    location: str
    hours: str | None = None
    contact: str | None = None
    method: str | None = Field(default="personal", description="Method of pickup: personal, mail, etc.")


class FoundItemCreate(BaseModel):
    """Schema for creating a new found item."""

    municipality: MunicipalityInfo
    item: ItemInfo
    pickup: PickupInfo
    categories: list[str] | None = None


class FoundItemUpdate(BaseModel):
    """Schema for updating an existing found item."""

    municipality: MunicipalityInfo | None = None
    item: ItemInfo | None = None
    pickup: PickupInfo | None = None
    categories: list[str] | None = None


class FoundItemResponse(BaseModel):
    """Response schema for a found item."""

    id: str
    municipality: MunicipalityInfo
    item: ItemInfo
    pickup: PickupInfo
    categories: list[str]
    createdAt: str | None = None
    updatedAt: str | None = None

    class Config:
        from_attributes = True
