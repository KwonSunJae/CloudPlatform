#!/usr/bin/python
#-*- coding: utf-8 -*-
import sh

def mountBindProcDevSyS(tmpDir):
    sh.mount("--bind","/proc",tmpDir+"/proc")
    sh.mount("--bind","/dev",tmpDir+"/dev")
    sh.mount("--bind","/sys",tmpDir+"/sys")
    
def unMountBindProcDevSyS(tmpDir):
    sh.umount(tmpDir+"/proc")
    sh.umount(tmpDir+"/dev")
    sh.umount(tmpDir+"/sys")
    sh.umount(tmpDir)
        
def unionmount(self, pathOfDiskProvFiles, directoryOfDiskProvFiles, pathOfDIskCreator):
    lowerDir=pathOfDIskCreator
    upperDir=sh.pathOfDiskProvFiles

