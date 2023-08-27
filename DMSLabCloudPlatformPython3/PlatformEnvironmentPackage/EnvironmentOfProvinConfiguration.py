#!/usr/bin/python
#-*- coding: utf-8 -*-
from Executor import mkDir
from DBRespositoryPackage import readConfigurationFilesToDB
from Utils import RedisDatabaseRepositoryData
import json

r = RedisDatabaseRepositoryData()


def initPathOfDiskImage(pathOfDiskImage):
    mkDir(pathOfDiskImage)
    
def initStorePathOfDiskImage(pathOfDiskImage):
    mkDir(pathOfDiskImage+"/diskImage")
    
def initProvinDirectoryOfDiskImage(pathOfDiskImage):
    upperdir=pathOfDiskImage+"/upperdir"
    mkDir(upperdir)
    mkDir(pathOfDiskImage+"/workdir")
    ProvinDirNameArrayList = json.loads(r.get_Hvalue("ConfigurationOfPlatform","ProvinDirNameArray"))
    for val in ProvinDirNameArrayList:
        mkDir(upperdir+"/"+val)
    
def initLowerDirectoryOfDiskImage(VDIUUID):
    mkDir("/tmp/"+VDIUUID)
    
def initEnvironmentOfDiskImage(VDIUUID):
    pathOfDiskImage= json.loads(r.get_Hvalue("ConfigurationOfPlatform","Path"))["diskProvisioningFiles"]+"/"+VDIUUID
    initPathOfDiskImage(pathOfDiskImage)
    initStorePathOfDiskImage(pathOfDiskImage)
    initProvinDirectoryOfDiskImage(pathOfDiskImage)
    initLowerDirectoryOfDiskImage(VDIUUID)

