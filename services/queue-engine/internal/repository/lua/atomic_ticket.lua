-- KEYS[1] = queue_key (Example: "queue:merchant-A:2026-01-28")
-- ARGV[1] = expiry_seconds (Example: 86400 for 24 hours)

local key = KEYS[1]
local ttl = tonumber(ARGV[1])

-- 1. Increment counter atomically
-- Redis guarantees this operation cannot be interrupted
local next_number = redis.call("INCR", key)

-- 2. If this is the first ticket (1), set expiration (TTL)
-- prevents garbage keys from accumulating forever
if next_number == 1 then
    redis.call("EXPIRE", key, ttl)
end

return next_number