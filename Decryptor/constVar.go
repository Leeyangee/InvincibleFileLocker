package main

//content VER1.1:ENC:AES1::
var AES1_HEADER = []byte{86, 69, 82, 49, 46, 49, 58, 69, 78, 67, 58, 65, 69, 83, 49, 58, 58}

//content VER1.1:ENC:RSA1::
var RSA1_HEADER = []byte{86, 69, 82, 49, 46, 49, 58, 69, 78, 67, 58, 82, 83, 65, 49, 58, 58}

var IS_DEBUG bool = false

var ENC_FILE_FIND int64 = 0
var ENC_FILE_DECRYPTED int64 = 0
