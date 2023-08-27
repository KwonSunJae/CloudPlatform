from Utils import execShellScript

GBToB=1024*1024*1024
OneMBToByte=1024*1024

def partedDefault(diskList, filesize):
    sizeToByte=filesize*GBToB
    partedSizeEnd=str(sizeToByte-1)+"B"
    partedSizeStart=str(OneMBToByte)+"B"
    execShellScript("parted -s "+diskList+" mklabel msdos")
    execShellScript("parted -s "+diskList+" mkpart primary \"ext3\" "+partedSizeStart+" "+partedSizeEnd)
    execShellScript("parted -s "+diskList+" set 1 boot on")
    
    primaryLoopDev=diskList+"p1"
    
    return primaryLoopDev