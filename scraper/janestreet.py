import requests

URL = "https://www.janestreet.com/jobs/main.json"

def fetch_janestreet_json():
    resp = requests.get(URL, timeout=10)
    resp.raise_for_status()
    return resp.json()

def main():
    data = fetch_janestreet_json()

    count = 1

    for d in data:
        print(d)

        if count == 5:
            break
        count += 1

    print(data)

if __name__ == "__main__":
    main()
