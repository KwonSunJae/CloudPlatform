#!/usr/bin/python
#-*- coding: utf-8 -*-
import datetime as dt

from VirtualDiskPackage import VirtualDiskImageBase
from DBRespositoryPackage import *
from Utils import SnowflakeIdGenerator
from BaseClass import Property
from Executor import *
from PlatformEnvironmentPackage import initEnvironmentOfDiskImage

r = RedisDatabaseRepositoryData()

#SNOWFLAKEINGENERATOR=SnowflakeIdGenerator(int(r.get_Hvalue("ConfigurationOfPlatform","CenterId")),int(r.get_Hvalue("ConfigurationOfPlatform","WorkerId")))
SNOWFLAKEINGENERATOR=SnowflakeIdGenerator(r.get_Hvalue("ConfigurationOfPlatform","CenterId"),r.get_Hvalue("ConfigurationOfPlatform","WorkerId"))
DATETIEMFORMAT="%Y-%m-%d %H:%M:%S"

class SystemSoftwareVirtualDiskImage(VirtualDiskImageBase):
    def __init__(self, format, size,softwareType, systemSoftwareInformationRefId, propertiesList):
        self.VDIUUID=str(SNOWFLAKEINGENERATOR.generate_id())
        super(SystemSoftwareVirtualDiskImage, self).__init__(self.VDIUUID, size, format, self.VDIUUID+"-diskImage", dt.datetime.now().strftime(DATETIEMFORMAT))
        self.systemSoftwareInformationRefId = systemSoftwareInformationRefId
        self.type=softwareType
        self.propertiesList = propertiesList
        r.insert_hash("SystemSoftwareDiskImages", self.VDIUUID, json.dumps(self.__dict__))
        
    
    def systemSoftwareVirtualDiskImageCreator(self):
       
       #init Environment of Creator SystemSoftware
#       initEnvironmentOfDiskImage(self.VDIUUID)
       print(r.exists_key("SystemSoftwareDiskImages"))
#       r.list_push("SystemSoftwareDiskImages",self.VDIUUID)
       

#       r.insert_hash(self.VDIUUID, "size", self.size)
#       r.insert_hash(self.VDIUUID, "createdTime", self.createdtime)
      
#      createDiskImage(self.format, self.size, self.VDIUUID, path+"/"+"diskImage")

    
#    def installingSystemSoftware(self):
        
#sysVDI = SystemSoftwareVirtualDiskImage(SnowflakeIdGenerator.generate_id(1207,202008917))

