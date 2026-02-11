"""DCAT-AP schemas for dane.gov.pl integration."""

from pydantic import Field
from typing import Optional, List

from .found_item import (
    MunicipalityInfo,
    ItemInfo,
    PickupInfo,
    FoundItemCreate,
    FoundItemUpdate,
    FoundItemResponse as BaseFoundItemResponse,
)


class FoundItemResponse(BaseFoundItemResponse):
    """Response model with dane.gov.pl / DCAT compatibility."""

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

    class Config:
        from_attributes = True
        populate_by_name = True
        json_schema_extra = {
            "example": {
                "id": "550e8400-e29b-41d4-a716-446655440000",
                "@context": "https://www.w3.org/ns/dcat",
                "@type": "dcat:Dataset",
                "dct:identifier": "550e8400-e29b-41d4-a716-446655440000",
                "dct:title": "Portfel skorzany brazowy",
                "dct:description": "Brazowy portfel ze skory naturalnej",
                "dct:license": "http://creativecommons.org/licenses/by/4.0/",
                "dcat:keyword": ["dokumenty", "portfele"],
                "municipality": {
                    "name": "Warszawa",
                    "type": "miasto",
                    "contactEmail": "kontakt@um.warszawa.pl"
                },
                "item": {
                    "name": "Portfel skorzany brazowy",
                    "category": "dokumenty",
                    "date": "2025-12-01",
                    "location": "Park Lazienkowski",
                    "status": "available",
                    "description": "Brazowy portfel ze skory"
                }
            }
        }
