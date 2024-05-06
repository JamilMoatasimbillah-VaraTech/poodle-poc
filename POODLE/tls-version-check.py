import requests

url = "https://localhost:9002"
response = requests.get(url)

# Check the TLS version
tls_version = response.connection.version
print("TLS version:", tls_version)
