"""Application configuration loaded from .env file."""

from pydantic_settings import BaseSettings
from typing import List


class Settings(BaseSettings):
    """Zguba.gov application settings."""

    DATABASE_URL: str = "sqlite+aiosqlite:///./zguba_gov.db"
    CORS_ORIGINS: str = "http://localhost:4200,http://localhost:3000"

    @property
    def cors_origins_list(self) -> List[str]:
        return [origin.strip() for origin in self.CORS_ORIGINS.split(",")]

    class Config:
        env_file = ".env"
        case_sensitive = True


settings = Settings()
