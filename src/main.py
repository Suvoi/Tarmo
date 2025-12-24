"""
Entry point for the FastAPI application.
Defines routes and dependency injection for database sessions.
"""

from fastapi import FastAPI, Depends, HTTPException, status
from src.database import Base, engine, SessionLocal

from src.recipes.controllers import router as recipes_router

from fastapi.middleware.cors import CORSMiddleware

# Automatically create all tables on startup (safe for dev/prototyping)
Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="Tarmo",
    version="2.0.0",
    description="Optimize and control batches based on recipes."
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(recipes_router, prefix="/recipes", tags=["Recipes"])