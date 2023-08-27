from Utils import execShellScript
import os
import sh

def createDiskImage(format, size, name, path):
    createtool="qemu-img"
    opt="create"
    diskImage=path+"/"+name+"."+format
    if (os.path.exists(diskImage)):
        print("File exsits start to remove same one ")
        sh.rm(diskImage)
    cmd_string=createtool+" "+opt+" "+diskImage+" "+size
    return execShellScript(cmd_string)

def makeFileSystemOfDiskImage(fileSystemType, diskImageMountPath):
    cmd_string="mkfs."+fileSystemType+" "+diskImageMountPath
    execShellScript(cmd_string)
    
def tune2fs(primaryLoopDev):
    cmd_string="tune2fs -c 0 -i 0 "+primaryLoopDev
    execShellScript(cmd_string)
    