import http from 'k6/http';
import { check, sleep } from 'k6';

// Config: 1000 Concurrent Users
export const options = {
  stages: [
    { duration: '30s', target: 50 },   // Ramp up to 50 users
    { duration: '1m', target: 200 },   // Ramp up to 200 users
    { duration: '30s', target: 500 },  // Spike to 500 users
    { duration: '1m', target: 1000 },  // Spike to 1000 users!
    { duration: '30s', target: 0 },    // Cooldown
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% of requests must be < 500ms
  },
};

const BASE_URL_QUEUE = 'http://localhost:3000';
const BASE_URL_INTELLIGENCE = 'http://localhost:8000';

export default function () {
  const merchantId = 'merchant-123';

  // 1. Take a Ticket (Go Service)
  const payload = JSON.stringify({
    merchant_id: merchantId,
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const resQueue = http.post(`${BASE_URL_QUEUE}/queue/take`, payload, params);

  check(resQueue, {
    'ticket created': (r) => r.status === 200,
    'has ticket number': (r) => r.json('ticket_number') !== undefined,
  });

  // 2. Check ETA (Python Service) - Simulated call by same user
  const resEta = http.get(`${BASE_URL_INTELLIGENCE}/eta/${merchantId}`);
  
  check(resEta, {
    'eta retrieved': (r) => r.status === 200,
  });

  sleep(1);
}
