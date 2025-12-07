from pydantic import BaseModel, EmailStr, Field
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
    method: Optional[str] = Field(default="personal", description="Method of pickup: personal, mail, etc.")


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
    """
    Response model z obsługą dane.gov.pl / DCAT compatibility
    """
    # DCAT Fields
    context: str = Field(default="https://www.w3.org/ns/dcat", alias="@context")
    type: str = Field(default="dcat:Dataset", alias="@type")
    dct_identifier: str = Field(alias="dct:identifier")
    dct_title: str = Field(alias="dct:title")
    dct_description: Optional[str] = Field(default=None, alias="dct:description")
    dct_issued: Optional[str] = Field(default=None, alias="dct:issued")
    dct_modified: Optional[str] = Field(default=None, alias="dct:modified")
    dct_license: str = Field(
        default="http://creativecommons.org/licenses/by/4.0/",
        alias="dct:license"
    )
    dcat_keyword: List[str] = Field(alias="dcat:keyword")
    dcat_landingPage: Optional[str] = Field(default=None, alias="dcat:landingPage")
    
    # Core Fields
    id: str
    municipality: MunicipalityInfo
    item: ItemInfo
    pickup: PickupInfo
    categories: List[str]
    createdAt: Optional[str] = None
    updatedAt: Optional[str] = None
    
    class Config:
        from_attributes = True
        populate_by_name = True  # Allow both original and aliased names
        json_schema_extra = {
            "example": {
                "id": "550e8400-e29b-41d4-a716-446655440000",
                "@context": "https://www.w3.org/ns/dcat",
                "@type": "dcat:Dataset",
                "dct:identifier": "550e8400-e29b-41d4-a716-446655440000",
                "dct:title": "Portfel skórzany brązowy",
                "dct:description": "Brązowy portfel ze skóry naturalnej",
                "dct:license": "http://creativecommons.org/licenses/by/4.0/",
                "dcat:keyword": ["dokumenty", "portfele"],
                "municipality": {
                    "name": "Warszawa",
                    "type": "miasto",
                    "contactEmail": "kontakt@um.warszawa.pl"
                },
                "item": {
                    "name": "Portfel skórzany brązowy",
                    "category": "dokumenty",
                    "date": "2025-12-01",
                    "location": "Park Łazienkowski",
                    "status": "available",
                    "description": "Brązowy portfel ze skóry"
                }
            }
        }
