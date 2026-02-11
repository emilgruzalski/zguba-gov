from contextlib import asynccontextmanager

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from database import init_db
from routers import found_items_router, metadata_router, odata_router, stats_router


@asynccontextmanager
async def lifespan(app: FastAPI):
    """Initialize database on startup"""
    await init_db()
    yield


app = FastAPI(
    title="Zguba.gov API",
    description="API dla systemu zgłaszania znalezionych rzeczy",
    version="1.0.0",
    lifespan=lifespan,
)

# CORS configuration - allow all localhost ports for development + dane.gov.pl
app.add_middleware(
    CORSMiddleware,
    allow_origins=[
        "http://localhost",
        "http://127.0.0.1",
        "https://dane.gov.pl",
        "https://api.dane.gov.pl",
        "https://data.europa.eu",
        "https://www.w3.org",
    ],
    allow_origin_regex=r"http://(localhost|127\.0\.0\.1)(:\d+)?",
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Include routers
app.include_router(metadata_router)
app.include_router(odata_router)
app.include_router(found_items_router)
app.include_router(stats_router)


@app.get("/")
async def root():
    return {"message": "Zguba.gov API", "version": "1.0.0", "docs": "/docs", "status": "running"}


@app.get("/health")
async def health_check():
    return {"status": "healthy"}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host="0.0.0.0", port=8000)
