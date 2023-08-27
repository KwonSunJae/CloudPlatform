class ProvioningVirtualDiskImageDTO:
    
    def __init__(self, diskSize, distribution, version, vmusername, vmpasswd, rootpwd):
        self.diskSize = diskSize
        self.distribution = distribution
        self.version = version
        self.codename = version
        self.vmusername = vmusername
        self.vmpasswd = vmpasswd
        self.rootpwd = rootpwd