from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)

def test_read_root():
    response = client.get("/")
    assert response.status_code  == 200
    assert response.json() == {"message": "Data Worker Ready 🧠"}

def test_get_eta():
    merchant_id = "TestResto"
    response = client.get(f"/eta/{merchant_id}")

    assert response.status_code == 200

    data = response.json()
    assert data["merchant_id"] == merchant_id
    assert "estimated_wait_time_minutes" in data
    "Ensure calculated average of mock data (10, 15, 12, 8, 20, 14, 11) is ~12.8, rounding up to 13."
    assert data["estimated_wait_time_minutes"] == 13