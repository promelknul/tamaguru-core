#!/usr/bin/env bash
set -euo pipefail
[ -z "${OZON_CLIENT_ID:-}" ] && [ -f .env ] && source .env
curl -s -o /tmp/resp.json -w '%{http_code}' \
  -X POST https://api-seller.ozon.ru/v3/posting/fbs/list \
  -H "Client-Id: ${OZON_CLIENT_ID}" \
  -H "Api-Key: ${OZON_API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"limit":1,"offset":0}' > /tmp/http_code
code=$(cat /tmp/http_code)
[ "$code" = "200" ] && jq -e 'has("result")' /tmp/resp.json >/dev/null && echo ok || echo fail
