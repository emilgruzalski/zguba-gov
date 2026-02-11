"""ORM model for found items."""

from sqlalchemy import Column, String, Integer, DateTime, Text, JSON
from datetime import datetime
from database import Base


class FoundItem(Base):
    """Found item database record.

    Stores item details, registering municipality, and pickup conditions.
    """

    __tablename__ = "found_items"

    id = Column(String, primary_key=True, index=True)

    # Municipality info
    municipality_name = Column(String, nullable=False, index=True)
    municipality_type = Column(String, nullable=False)
    municipality_email = Column(String, nullable=False)

    # Item info
    item_name = Column(String, nullable=False, index=True)
    item_category = Column(String, nullable=False, index=True)
    item_date = Column(String, nullable=False)
    item_location = Column(String, nullable=False)
    item_status = Column(String, nullable=False, default="available")
    item_description = Column(Text, nullable=True)

    # Pickup info
    pickup_deadline = Column(Integer, nullable=False)
    pickup_location = Column(String, nullable=False)
    pickup_hours = Column(String, nullable=True)
    pickup_contact = Column(String, nullable=True)

    # Additional fields
    categories = Column(JSON, nullable=True)
    created_at = Column(DateTime, default=datetime.utcnow, index=True)
    updated_at = Column(DateTime, default=datetime.utcnow, onupdate=datetime.utcnow)

    def to_dict(self):
        """Convert record to a dictionary matching the API schema."""
        return {
            "id": self.id,
            "municipality": {
                "name": self.municipality_name,
                "type": self.municipality_type,
                "contactEmail": self.municipality_email
            },
            "item": {
                "name": self.item_name,
                "category": self.item_category,
                "date": self.item_date,
                "location": self.item_location,
                "status": self.item_status,
                "description": self.item_description
            },
            "pickup": {
                "deadline": self.pickup_deadline,
                "location": self.pickup_location,
                "hours": self.pickup_hours,
                "contact": self.pickup_contact
            },
            "categories": self.categories or [],
            "createdAt": self.created_at.isoformat() if self.created_at else None,
            "updatedAt": self.updated_at.isoformat() if self.updated_at else None
        }
