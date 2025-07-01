#!/usr/bin/env bash
set -euo pipefail
r=$(curl -s -w '%{http_code}' -o /tmp/resp.json \
  -X POST https://api-seller.ozon.ru/v3/posting/fbs/list \
  -H "Client-Id: ${OZON_CLIENT_ID}" \
  -H "Api-Key: ${OZON_API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"limit":1,"offset":0}')
[ "$r" = "200" ] && jq -e 'has("result")' /tmp/resp.json >/dev/null && echo ok || echo fail
