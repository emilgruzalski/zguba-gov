"""
OData Router dla dane.gov.pl compatibility
"""
from fastapi import APIRouter, Depends, Query
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, desc
from typing import Optional
from database import get_db
from models import FoundItem
from schemas import FoundItemResponse

router = APIRouter(prefix="/odata", tags=["odata"])


@router.get("", response_model=dict)
async def get_odata_items(
    skip: int = Query(0, alias="$skip", ge=0),
    limit: int = Query(50, alias="$top", ge=1, le=100),
    filter: Optional[str] = Query(None, alias="$filter", description="OData filter expression"),
    select_fields: Optional[str] = Query(None, alias="$select", description="OData select expression"),
    orderby: Optional[str] = Query(None, alias="$orderby", description="OData orderby expression"),
    count: bool = Query(False, alias="$count", description="Include count in response"),
    db: AsyncSession = Depends(get_db)
):
    """
    OData endpoint dla dane.gov.pl compatibility
    
    Supported Query Options:
    - $skip=N - Skip N records
    - $top=N - Take N records
    - $filter=expression - Filter results
    - $select=field1,field2 - Select specific fields
    - $orderby=field asc|desc - Order results
    - $count=true - Include total count
    
    Examples:
    - /odata?$filter=item_status eq 'available'&$top=10
    - /odata?$skip=5&$top=20&$orderby=created_at desc
    - /odata?$select=municipality_name,item_name
    - /odata?$count=true
    """
    
    query = select(FoundItem)
    
    # Apply Filter
    if filter:
        query = _apply_odata_filter(query, filter)
    
    # Apply OrderBy
    if orderby:
        query = _apply_odata_orderby(query, orderby)
    
    # Apply Pagination
    query = query.offset(skip).limit(limit)
    
    # Execute Query
    result = await db.execute(query)
    items = result.scalars().all()
    
    # Format Response
    items_data = [item.to_dict() for item in items]
    
    response = {
        "odata.context": "https://api.zguba.gov/odata/$metadata",
        "value": items_data
    }
    
    if count:
        count_result = await db.execute(select(FoundItem))
        total_count = len(count_result.scalars().all())
        response["odata.count"] = total_count
    
    return response


@router.get("/$metadata", response_model=dict)
async def get_odata_metadata():
    """
    OData Metadata - zwraca schemat dostępnych pól
    """
    return {
        "edmx:Edmx": {
            "@xmlns:edmx": "http://schemas.microsoft.com/ado/2007/06/edmx",
            "@Version": "1.0",
            "edmx:DataServices": {
                "@m:DataServiceVersion": "2.0",
                "Schema": {
                    "@xmlns": "http://schemas.microsoft.com/ado/2008/09/edm",
                    "@Namespace": "Zguba.Models",
                    "EntityType": {
                        "@Name": "FoundItem",
                        "Key": {
                            "PropertyRef": {
                                "@Name": "id"
                            }
                        },
                        "Property": [
                            {"@Name": "id", "@Type": "Edm.String", "@Nullable": "false"},
                            {"@Name": "municipality_name", "@Type": "Edm.String"},
                            {"@Name": "municipality_type", "@Type": "Edm.String"},
                            {"@Name": "municipality_email", "@Type": "Edm.String"},
                            {"@Name": "item_name", "@Type": "Edm.String"},
                            {"@Name": "item_category", "@Type": "Edm.String"},
                            {"@Name": "item_date", "@Type": "Edm.String"},
                            {"@Name": "item_location", "@Type": "Edm.String"},
                            {"@Name": "item_status", "@Type": "Edm.String"},
                            {"@Name": "item_description", "@Type": "Edm.String"},
                            {"@Name": "pickup_deadline", "@Type": "Edm.Int32"},
                            {"@Name": "pickup_location", "@Type": "Edm.String"},
                            {"@Name": "pickup_hours", "@Type": "Edm.String"},
                            {"@Name": "pickup_contact", "@Type": "Edm.String"},
                            {"@Name": "created_at", "@Type": "Edm.DateTime"},
                            {"@Name": "updated_at", "@Type": "Edm.DateTime"}
                        ]
                    },
                    "EntityContainer": {
                        "@Name": "ZgubaService",
                        "@m:IsDefaultEntityContainer": "true",
                        "EntitySet": {
                            "@Name": "FoundItems",
                            "@EntityType": "Zguba.Models.FoundItem"
                        }
                    }
                }
            }
        }
    }


def _apply_odata_filter(query, filter_expr: str):
    """
    Parse and apply OData filter expressions
    
    Supported:
    - field eq 'value'
    - field ne 'value'
    - field gt 123
    - field lt 123
    - startswith(field, 'prefix')
    - contains(field, 'substring')
    """
    from sqlalchemy import and_, or_
    from models import FoundItem
    
    # Simple parsing - can be expanded with proper OData parser
    
    # eq operator
    if " eq " in filter_expr:
        parts = filter_expr.split(" eq ")
        field_name = parts[0].strip()
        field_value = parts[1].strip().strip("'\"")
        
        if field_name == "item_status":
            query = query.where(FoundItem.item_status == field_value)
        elif field_name == "item_category":
            query = query.where(FoundItem.item_category == field_value)
        elif field_name == "municipality_name":
            query = query.where(FoundItem.municipality_name == field_value)
    
    # contains operator
    elif "contains(" in filter_expr:
        import re
        match = re.search(r"contains\((\w+),\s*'(.+?)'\)", filter_expr)
        if match:
            field_name, search_value = match.groups()
            if field_name == "item_name":
                query = query.where(FoundItem.item_name.ilike(f"%{search_value}%"))
            elif field_name == "item_description":
                query = query.where(FoundItem.item_description.ilike(f"%{search_value}%"))
    
    # startswith operator
    elif "startswith(" in filter_expr:
        import re
        match = re.search(r"startswith\((\w+),\s*'(.+?)'\)", filter_expr)
        if match:
            field_name, prefix_value = match.groups()
            if field_name == "municipality_name":
                query = query.where(FoundItem.municipality_name.ilike(f"{prefix_value}%"))
    
    return query


def _apply_odata_orderby(query, orderby_expr: str):
    """
    Parse and apply OData orderby expressions
    
    Format: field asc|desc
    Examples: created_at desc, item_name asc
    """
    from models import FoundItem
    
    parts = orderby_expr.split()
    field_name = parts[0].strip()
    direction = "asc"
    
    if len(parts) > 1:
        direction = parts[1].lower()
    
    if field_name == "created_at":
        if direction == "desc":
            query = query.order_by(FoundItem.created_at.desc())
        else:
            query = query.order_by(FoundItem.created_at)
    elif field_name == "item_name":
        if direction == "desc":
            query = query.order_by(FoundItem.item_name.desc())
        else:
            query = query.order_by(FoundItem.item_name)
    elif field_name == "item_date":
        if direction == "desc":
            query = query.order_by(FoundItem.item_date.desc())
        else:
            query = query.order_by(FoundItem.item_date)
    
    return query
