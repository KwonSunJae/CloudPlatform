#!/usr/bin/python
#-*- coding: utf-8 -*-
from FlaskRestulFulAPI.v1.SystemSoftwareInformation.SystemSoftwareInformationRepository import SystemSoftwareInformationOperation
from FlaskRestulFulAPI.v1.SystemSoftwareInformation.SystemSoftwareInformationDAO import SystemSoftwareInformation
#from DBRespositoryPackage import *
from Utils import readYAMLFileToDictType
# format, size,softwareType, systemSoftwareInformationRefId, propertiesList

#sysSWVDI=SystemSoftwareVirtualDiskImage("qcow2","120G","systemsoftware",3, None)
#sysSWVDI.systemSoftwareVirtualDiskImageCreator()


#readConfigurationFilesToDB()
# Connection for Database

# Read sysSWInfo data from YAML file
dbdata=readYAMLFileToDictType("Documents/SoftwareInfo.yaml","r")
systemsoftwaredata=dbdata["SystemSoftware"]
sysSWOperation = SystemSoftwareInformationOperation()
#for val in systemsoftwaredata:
#    for arch in val["architecture"]:
#    # def __init__(self, sysSWId, distribution, version,codename, architecture, description):
#       sysSWDAO=SystemSoftwareInformation(distribution=val["distribution"], version=val["version"], codename=val["codename"], architecture=arch, description=val["description"])
#       sysSWOperation.insertSingleSystemSoftwareInformatioDAO(sysSWDAO)
sysSWDAO=SystemSoftwareInformation(distribution="Ubuntu", version="20.04", codename="focal", architecture="amd64", description="")

sysSWOperation.dataExisted(sysSWDAO)
