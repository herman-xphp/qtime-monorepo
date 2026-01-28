from sqlalchemy import create_engine
from app.core.config import settings

# Create Engine
engine = create_engine(settings.DATABASE_URL)
