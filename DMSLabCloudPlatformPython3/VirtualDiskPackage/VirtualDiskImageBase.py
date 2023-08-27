#!/usr/bin/python
#-*- coding: utf-8 -*-
from Executor import *


class VirtualDiskImageBase:
    def __init__(self, VDIUUID, size, format, name, createdtime):
        self.VDIUUID = VDIUUID
        self.size = size
        self.format = format
        self.name = name
        self.createdtime = createdtime
