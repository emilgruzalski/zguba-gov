"""Metadata router for dane.gov.pl integration."""
from fastapi import APIRouter
from datetime import datetime

router = APIRouter(prefix="/metadata", tags=["metadata"])


@router.get("", tags=["metadata"])
async def get_dataset_metadata():
    """Return dataset metadata in DCAT-AP format, compatible with dane.gov.pl."""
    return {
        "@context": "https://www.w3.org/ns/dcat",
        "@type": "dcat:Catalog",
        "dct:title": "Katalog Rzeczy Znalezionych - Zguba.gov",
        "dct:description": "System zarabiania znalezionych przedmiotów w jednostkach administracji publicznej w Polsce",
        "dct:issued": "2025-12-01T00:00:00Z",
        "dct:modified": datetime.utcnow().isoformat() + "Z",
        "dct:language": "pl",
        "dcat:homepage": "https://zguba.gov",
        
        "dcat:dataset": [
            {
                "@type": "dcat:Dataset",
                "dct:identifier": "found-items-dataset",
                "dct:title": "Znalezione Rzeczy",
                "dct:description": "Bieżąca lista znalezionych przedmiotów w jednostkach administracji publicznej",
                "dct:issued": "2025-12-01T00:00:00Z",
                "dct:modified": datetime.utcnow().isoformat() + "Z",
                "dct:language": "pl",
                "dct:license": "http://creativecommons.org/licenses/by/4.0/",
                "dcat:keyword": ["rzeczy znalezione", "przedmioty", "administracja publiczna", "dane otwarte"],
                "dcat:theme": ["http://publications.europa.eu/resource/authority/data-theme/SOCI"],
                "dct:accrualPeriodicity": "http://purl.org/ckan/freq/daily",
                
                "dcat:distribution": [
                    {
                        "@type": "dcat:Distribution",
                        "dct:title": "JSON API",
                        "dct:description": "REST API zwracający dane w formacie JSON",
                        "dcat:accessURL": "https://api.zguba.gov/api/found-items",
                        "dcat:downloadURL": "https://api.zguba.gov/api/found-items?limit=1000",
                        "dct:format": "JSON",
                        "dcat:mediaType": "application/json",
                        "dct:license": "http://creativecommons.org/licenses/by/4.0/"
                    },
                    {
                        "@type": "dcat:Distribution",
                        "dct:title": "OData API",
                        "dct:description": "OData endpoint dla zaawansowanego filtrowania",
                        "dcat:accessURL": "https://api.zguba.gov/odata",
                        "dct:format": "OData",
                        "dcat:mediaType": "application/json"
                    }
                ],
                
                "dcat:contactPoint": {
                    "@type": "vcard:Organization",
                    "vcard:fn": "Zguba.gov Support",
                    "vcard:hasEmail": "mailto:support@zguba.gov"
                }
            }
        ],
        
        "dcat:organization": {
            "@type": "foaf:Organization",
            "foaf:name": "Zguba.gov"
        }
    }


@router.get("/dcat", tags=["metadata"])
async def get_dcat_rdf():
    """Return metadata in Turtle (RDF) format."""
    return {
        "message": "RDF endpoint - TODO",
        "format": "text/turtle",
        "uri": "https://api.zguba.gov/metadata/dcat.rdf"
    }


@router.get("/distribution/{distribution_id}", tags=["metadata"])
async def get_distribution_metadata(distribution_id: str):
    """Return metadata for a specific distribution."""
    distributions = {
        "json-api": {
            "@type": "dcat:Distribution",
            "dct:title": "JSON API",
            "dcat:accessURL": "https://api.zguba.gov/api/found-items",
            "dct:format": "JSON",
            "dcat:mediaType": "application/json",
            "dct:license": "http://creativecommons.org/licenses/by/4.0/"
        },
        "odata": {
            "@type": "dcat:Distribution",
            "dct:title": "OData API",
            "dcat:accessURL": "https://api.zguba.gov/odata",
            "dct:format": "OData",
            "dcat:mediaType": "application/json"
        }
    }
    
    return distributions.get(distribution_id, {"error": "Distribution not found"})
