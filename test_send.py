import socket

sock = socket.socket()
sock.connect(('localhost', 4040))
sock.send(b'press;t')
