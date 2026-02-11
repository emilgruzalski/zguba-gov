from fastapi import APIRouter, Depends
from sqlalchemy import func, select
from sqlalchemy.ext.asyncio import AsyncSession

from database import get_db
from models import FoundItem

router = APIRouter(prefix="/api/stats", tags=["stats"])


@router.get("")
async def get_stats(db: AsyncSession = Depends(get_db)):
    """Get overall system statistics."""
    # Count total found items
    total_items_result = await db.execute(select(func.count(FoundItem.id)))
    total_items = total_items_result.scalar()

    # Count by status
    available_result = await db.execute(select(func.count(FoundItem.id)).where(FoundItem.item_status == "available"))
    available_items = available_result.scalar()

    claimed_result = await db.execute(select(func.count(FoundItem.id)).where(FoundItem.item_status == "claimed"))
    claimed_items = claimed_result.scalar()

    # Count by category
    categories_result = await db.execute(
        select(FoundItem.item_category, func.count(FoundItem.id).label("count"))
        .group_by(FoundItem.item_category)
        .order_by(func.count(FoundItem.id).desc())
        .limit(10)
    )
    categories = [{"category": row[0], "count": row[1]} for row in categories_result.all()]

    # Top municipalities
    municipalities_result = await db.execute(
        select(FoundItem.municipality_name, func.count(FoundItem.id).label("count"))
        .group_by(FoundItem.municipality_name)
        .order_by(func.count(FoundItem.id).desc())
        .limit(10)
    )
    top_municipalities = [{"name": row[0], "count": row[1]} for row in municipalities_result.all()]

    return {
        "foundItems": {"total": total_items, "available": available_items, "claimed": claimed_items},
        "topCategories": categories,
        "topMunicipalities": top_municipalities,
    }
