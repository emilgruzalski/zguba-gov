"""Pydantic schemas for found items."""

from pydantic import BaseModel, EmailStr, Field
from typing import Optional, List


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
    description: Optional[str] = None


class PickupInfo(BaseModel):
    """Item pickup conditions."""
    deadline: int
    location: str
    hours: Optional[str] = None
    contact: Optional[str] = None
    method: Optional[str] = Field(
        default="personal",
        description="Method of pickup: personal, mail, etc."
    )


class FoundItemCreate(BaseModel):
    """Schema for creating a new found item."""
    municipality: MunicipalityInfo
    item: ItemInfo
    pickup: PickupInfo
    categories: Optional[List[str]] = None


class FoundItemUpdate(BaseModel):
    """Schema for updating an existing found item."""
    municipality: Optional[MunicipalityInfo] = None
    item: Optional[ItemInfo] = None
    pickup: Optional[PickupInfo] = None
    categories: Optional[List[str]] = None


class FoundItemResponse(BaseModel):
    """Response schema for a found item."""
    id: str
    municipality: MunicipalityInfo
    item: ItemInfo
    pickup: PickupInfo
    categories: List[str]
    createdAt: Optional[str] = None
    updatedAt: Optional[str] = None

    class Config:
        from_attributes = True
