import sh
import os
from BaseClass import Message

def mkTempDir(virtualDiskImageName):
    virtualDiskImageTempDir=virtualDiskImageName+'.XXXXXX'
    return sh.mktemp("-d","-t",virtualDiskImageTempDir)

def mkDir(dirIncludingPath):
    message=Message()
    if not os.path.exists(dirIncludingPath):
        os.makedirs(dirIncludingPath)
        message.status(True)
    else:
        message.status(False)
        message.msg(dirIncludingPath+" directory already exists error")
    return message