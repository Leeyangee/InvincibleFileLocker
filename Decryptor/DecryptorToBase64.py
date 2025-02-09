import base64
import pyperclip as cb
import os

os.system('''go build -ldflags "-H=windowsgui" .\\main.go .\\linkedVar.go .\\encrypt.go .\\dirOperaDecryptor.go .\\file.go .\\constVar.go''')

with open('main.exe', 'rb') as f:
    cb.copy(base64.b64encode(f.read()).decode())

os.system('''del main.exe''')