import re
from Utils import execShellScript



def losetupFunction(pathOfVirtualDiskImage):
    cmd_string="losetup -P -f --show "+pathOfVirtualDiskImage
    result=execShellScript(cmd_string)
    if (result.status):
        loop=re.sub(r'/dev/', '',result.value, count=1)
        diskList="/dev/"+loop
    
    