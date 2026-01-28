from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    DATABASE_URL: str = "postgresql://postgres@localhost:5432/qtime_db"

    class Config:
        env_file = ".env"

settings = Settings()
