import requests
try:
    r=requests.get("http://127.0.0.1:8888/ping")
except:
    print("error")
else:   
    print(r.text)
try:
    r=requests.get("http://127.0.0.1:8888/echo?message=Hello")
except:
    print("error")
else:
    print(r.text)
