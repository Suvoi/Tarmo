"""
Entry point for the FastAPI application.
Defines routes and dependency injection for database sessions.
"""

from fastapi import FastAPI, Depends, HTTPException, status
from src.database import Base, engine, SessionLocal
from src.recipes.routes import router as recipes_router

from fastapi.middleware.cors import CORSMiddleware

# Automatically create all tables on startup (safe for dev/prototyping)
Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="Tarmo",
    version="0.0.1",
    description="Optimize and control batches with recipes."
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(recipes_router, prefix="/recipes", tags=["Recipes"])