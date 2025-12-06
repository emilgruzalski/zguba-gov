from fastapi import APIRouter, Depends, Query, HTTPException
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, or_, func, desc
from typing import List, Optional
from datetime import datetime
import uuid
from database import get_db
from models import FoundItem
from schemas import FoundItemCreate, FoundItemUpdate, FoundItemResponse

router = APIRouter(prefix="/api/found-items", tags=["found-items"])


@router.get("", response_model=List[FoundItemResponse])
async def get_found_items(
    skip: int = Query(0, ge=0),
    limit: int = Query(50, ge=1, le=100),
    category: Optional[str] = Query(None, description="Kategoria przedmiotu"),
    municipality: Optional[str] = Query(None, description="Nazwa gminy"),
    status: Optional[str] = Query(None, description="Status: available, claimed, expired"),
    search: Optional[str] = Query(None, description="Wyszukaj w nazwie lub opisie"),
    db: AsyncSession = Depends(get_db)
):
    """
    Pobierz listę znalezionych rzeczy z opcjonalnymi filtrami
    """
    query = select(FoundItem)
    
    if category:
        query = query.where(FoundItem.item_category == category)
    
    if municipality:
        query = query.where(FoundItem.municipality_name.ilike(f"%{municipality}%"))
    
    if status:
        query = query.where(FoundItem.item_status == status)
    
    if search:
        search_pattern = f"%{search}%"
        query = query.where(
            or_(
                FoundItem.item_name.ilike(search_pattern),
                FoundItem.item_description.ilike(search_pattern),
                FoundItem.item_location.ilike(search_pattern)
            )
        )
    
    query = query.order_by(desc(FoundItem.created_at)).offset(skip).limit(limit)
    
    result = await db.execute(query)
    items = result.scalars().all()
    
    return [FoundItemResponse(**item.to_dict()) for item in items]


@router.post("", response_model=FoundItemResponse, status_code=201)
async def create_found_item(
    item_data: FoundItemCreate,
    db: AsyncSession = Depends(get_db)
):
    """
    Dodaj nową znalezioną rzecz
    """
    item_id = str(uuid.uuid4())
    
    db_item = FoundItem(
        id=item_id,
        municipality_name=item_data.municipality.name,
        municipality_type=item_data.municipality.type,
        municipality_email=item_data.municipality.contactEmail,
        item_name=item_data.item.name,
        item_category=item_data.item.category,
        item_date=item_data.item.date,
        item_location=item_data.item.location,
        item_status=item_data.item.status,
        item_description=item_data.item.description,
        pickup_deadline=item_data.pickup.deadline,
        pickup_location=item_data.pickup.location,
        pickup_hours=item_data.pickup.hours,
        pickup_contact=item_data.pickup.contact,
        categories=item_data.categories or []
    )
    
    db.add(db_item)
    await db.commit()
    await db.refresh(db_item)
    
    return FoundItemResponse(**db_item.to_dict())


@router.get("/{item_id}", response_model=FoundItemResponse)
async def get_found_item(
    item_id: str,
    db: AsyncSession = Depends(get_db)
):
    """
    Pobierz szczegóły znalezionej rzeczy
    """
    result = await db.execute(
        select(FoundItem).where(FoundItem.id == item_id)
    )
    item = result.scalar_one_or_none()
    
    if not item:
        raise HTTPException(status_code=404, detail="Przedmiot nie znaleziony")
    
    return FoundItemResponse(**item.to_dict())


@router.put("/{item_id}", response_model=FoundItemResponse)
async def update_found_item(
    item_id: str,
    item_data: FoundItemUpdate,
    db: AsyncSession = Depends(get_db)
):
    """
    Zaktualizuj znalezioną rzecz
    """
    result = await db.execute(
        select(FoundItem).where(FoundItem.id == item_id)
    )
    db_item = result.scalar_one_or_none()
    
    if not db_item:
        raise HTTPException(status_code=404, detail="Przedmiot nie znaleziony")
    
    # Update fields
    if item_data.municipality:
        db_item.municipality_name = item_data.municipality.name
        db_item.municipality_type = item_data.municipality.type
        db_item.municipality_email = item_data.municipality.contactEmail
    
    if item_data.item:
        if item_data.item.name:
            db_item.item_name = item_data.item.name
        if item_data.item.category:
            db_item.item_category = item_data.item.category
        if item_data.item.date:
            db_item.item_date = item_data.item.date
        if item_data.item.location:
            db_item.item_location = item_data.item.location
        if item_data.item.status:
            db_item.item_status = item_data.item.status
        if item_data.item.description is not None:
            db_item.item_description = item_data.item.description
    
    if item_data.pickup:
        if item_data.pickup.deadline:
            db_item.pickup_deadline = item_data.pickup.deadline
        if item_data.pickup.location:
            db_item.pickup_location = item_data.pickup.location
        if item_data.pickup.hours is not None:
            db_item.pickup_hours = item_data.pickup.hours
        if item_data.pickup.contact is not None:
            db_item.pickup_contact = item_data.pickup.contact
    
    if item_data.categories is not None:
        db_item.categories = item_data.categories
    
    db_item.updated_at = datetime.utcnow()
    
    await db.commit()
    await db.refresh(db_item)
    
    return FoundItemResponse(**db_item.to_dict())


@router.delete("/{item_id}", status_code=204)
async def delete_found_item(
    item_id: str,
    db: AsyncSession = Depends(get_db)
):
    """
    Usuń znalezioną rzecz
    """
    result = await db.execute(
        select(FoundItem).where(FoundItem.id == item_id)
    )
    db_item = result.scalar_one_or_none()
    
    if not db_item:
        raise HTTPException(status_code=404, detail="Przedmiot nie znaleziony")
    
    await db.delete(db_item)
    await db.commit()
    
    return None


@router.get("/categories/list")
async def get_categories(db: AsyncSession = Depends(get_db)):
    """
    Pobierz listę wszystkich używanych kategorii
    """
    result = await db.execute(
        select(FoundItem.item_category)
        .distinct()
        .order_by(FoundItem.item_category)
    )
    categories = result.scalars().all()
    
    return [{"value": cat, "label": cat.capitalize()} for cat in categories if cat]
