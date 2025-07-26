import requests, logging, time, os

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")

HEAD = {
    "Client-Id": os.getenv("OZON_CLIENT_ID"),
    "Api-Key": os.getenv("OZON_API_KEY"),
    "Content-Type": "application/json"
}
def fetch():
    body = {
    "dir": "ASC",
    "limit": 50,
    "offset": 0,
    "with": {"analytics_data": True, "barcodes": True},
    "filter": {
        "since": (dt.datetime.utcnow()-dt.timedelta(days=1)).isoformat()+"Z",
        "to": dt.datetime.utcnow().isoformat()+"Z",
        "status": "awaiting_packaging"
    }
}
}
    r = requests.post("https://api-seller.ozon.ru/v3/posting/fbs/list",
                      headers=HEAD, json=body, timeout=30)
    r.raise_for_status()
    return r.json()["result"]["postings"]

def fetch():
    body = {
    "dir": "ASC",
    "limit": 50,
    "offset": 0,
    "with": {"analytics_data": True, "barcodes": True},
    "filter": {
        "since": (dt.datetime.utcnow()-dt.timedelta(days=1)).isoformat()+"Z",
        "to": dt.datetime.utcnow().isoformat()+"Z",
        "status": "awaiting_packaging"
    }
}
        },
        "limit": 5,
        "offset": 0
    }
    r = requests.post("https://api-seller.ozon.ru/v3/posting/fbs/list",
                      headers=HEAD, json=body, timeout=30)
    r.raise_for_status()
    return r.json()["result"]["postings"]

while True:
    try:
        postings = fetch()
        logging.info("Fetched %s orders", len(postings))
    except Exception as e:
        logging.exception("ERR %s", e)
    time.sleep(900)
import requests, logging, time, os

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")

HEAD = {
    "Client-Id": os.getenv("OZON_CLIENT_ID"),
    "Api-Key": os.getenv("OZON_API_KEY"),
    "Content-Type": "application/json"
}

def fetch():
    body = {
    "dir": "ASC",
    "limit": 50,
    "offset": 0,
    "with": {"analytics_data": True, "barcodes": True},
    "filter": {
        "since": (dt.datetime.utcnow()-dt.timedelta(days=1)).isoformat()+"Z",
        "to": dt.datetime.utcnow().isoformat()+"Z",
        "status": "awaiting_packaging"
    }
}
        },
        "limit": 5,
        "offset": 0
    }
    r = requests.post("https://api-seller.ozon.ru/v3/posting/fbs/list",
                      headers=HEAD, json=body, timeout=30)
    r.raise_for_status()
    return r.json()["result"]["postings"]

while True:
    try:
        postings = fetch()
        logging.info("Fetched %s orders", len(postings))
    except Exception as e:
        logging.exception("ERR %s", e)
    time.sleep(900)
