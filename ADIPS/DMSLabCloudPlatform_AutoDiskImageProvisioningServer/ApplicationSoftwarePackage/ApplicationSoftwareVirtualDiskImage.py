#!/usr/bin/python
#-*- coding: utf-8 -*-

from VirtualDiskPackage.VirtualDiskImageBase import VirtualDiskImageBase

class ApplicationSoftwareVirtualDiskImage(VirtualDiskImageBase):
    def __init__(self, applicationSoftwareInfornationRefId, appSoftwareVDIMetaData, propertyList):
        self.applicationSoftwareInfornationRefId = None
        self.apppSoftwareVDIMetaData = None
        self.propertyList = []

