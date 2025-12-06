from pydantic import BaseModel, EmailStr
from typing import Optional, List
from datetime import datetime


class MunicipalityInfo(BaseModel):
    name: str
    type: str
    contactEmail: EmailStr


class ItemInfo(BaseModel):
    name: str
    category: str
    date: str
    location: str
    status: str = "available"
    description: Optional[str] = None


class PickupInfo(BaseModel):
    deadline: int
    location: str
    hours: Optional[str] = None
    contact: Optional[str] = None


class FoundItemCreate(BaseModel):
    municipality: MunicipalityInfo
    item: ItemInfo
    pickup: PickupInfo
    categories: Optional[List[str]] = None


class FoundItemUpdate(BaseModel):
    municipality: Optional[MunicipalityInfo] = None
    item: Optional[ItemInfo] = None
    pickup: Optional[PickupInfo] = None
    categories: Optional[List[str]] = None


class FoundItemResponse(BaseModel):
    id: str
    municipality: MunicipalityInfo
    item: ItemInfo
    pickup: PickupInfo
    categories: List[str]
    createdAt: Optional[str] = None
    updatedAt: Optional[str] = None
    
    class Config:
        from_attributes = True
