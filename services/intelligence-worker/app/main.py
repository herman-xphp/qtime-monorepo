import pandas as pd
from fastapi import FastAPI
from app.db.session import engine

app = FastAPI(title="Q-Time Intelligence")

@app.get("/")
def read_root():
    return {"message": "Data Worker Ready 🧠"}

@app.get("/eta/{merchant_id}")
def get_eta(merchant_id: str):
    # 1. Data Ingestion
    # Imagine querying the Postgres Database:
    # query = f"SELECT duration_minutes FROM queue_logs WHERE merchant_id = '{merchant_id}'"
    # df = pd.read_sql(query, engine)

    # For Demo pursoses, we mock historical queue data (since table queue_logs doesn't exist yet)
    # These numbers represent previous service durations (in minutes)
    mock_data = {
        "duration_minutes": [10, 15, 12, 8, 20, 14, 11]
    }
    df = pd.DataFrame(mock_data)

    # 2. Intelligent Processing
    # Calculate Mean (Average)
    avg_time = df["duration_minutes"].mean()

    # We could add more logic here: Median, P95, etc.

    return {
        "merchant_id": merchant_id,
        "estimated_wait_time_minutes": round(avg_time), # Expected Output: 13
        "samples_count": len(df),
        "algorithm": "Simple Moving Average (Pandas)"
    }