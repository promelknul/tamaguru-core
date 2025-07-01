#!/usr/bin/env bash
set -euo pipefail
curl -s -X POST https://api-seller.ozon.ru/v3/posting/fbs/list \
  -H "Client-Id: ${OZON_CLIENT_ID}" \
  -H "Api-Key: ${OZON_API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"limit":1,"offset":0}' |
  jq -e '.result | length > 0' >/dev/null && echo ok || echo fail
