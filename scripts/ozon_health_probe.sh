#!/usr/bin/env bash
set -euo pipefail
[ -z "${OZON_CLIENT_ID:-}" ] && [ -f .env ] && source .env
payload='{"dir":"ASC","filter":{"processed_at_from":"2024-01-01T00:00:00Z","processed_at_to":"2024-12-31T23:59:59Z"},"limit":1,"offset":0}'
read -r body code <<<"$(curl -s -w ' %{http_code}' -X POST https://api-seller.ozon.ru/v3/posting/fbs/list \
  -H "Client-Id: ${OZON_CLIENT_ID}" \
  -H "Api-Key: ${OZON_API_KEY}" \
  -H "Content-Type: application/json" \
  -d "$payload")"
json="${body% ${code}}"
if [ "$code" = "200" ]; then
  echo ok
else
  echo "fail ($code)"; echo "$json" | jq .
  exit 1
fi
