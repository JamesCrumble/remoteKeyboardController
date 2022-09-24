import os

SERVER_PORT: int = 7575
assert 1025 <= SERVER_PORT <= 60000

os.system('pip install json-server.py')
os.system('pip install requests')
os.system("json-server settings.json -b %d" % (SERVER_PORT,))
