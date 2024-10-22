from urllib.error import URLError
from urllib.request import urlopen


def check_url(url):
    try:
        with urlopen(url, timeout=5):
            return True
    except (URLError, TimeoutError):
        return False


url = 'https://www.linkedin.com/learning'
print(check_url(url))
