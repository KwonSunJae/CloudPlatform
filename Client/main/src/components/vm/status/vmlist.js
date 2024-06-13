// VmList.js
import React from 'react';
import instance from '../../../apis/instance'

const VmList = ({ data }) => {
    const handleDelete = (vmName) => {
        // Add logic for deleting the VM with the given name
        console.log(`Delete VM: ${vmName}`);
        instance.get("/vm")
        .then((response)=>{
            var datas = response.data.data
            var piv ="";
            console.log(datas);
            for(var i =0; i<datas.length; i++){
                if(vmName=== datas[i].Name){
                    piv=datas[i].Id;
                    break;
                }
            }
            if(piv === ""){
                console.error("no vms");
            }
            else{
                instance.delete("/vm/"+piv)
                .then((response)=>{
                    console.log(response);
                    window.location.reload();
                })
                .catch((err)=>{
                    console.log(err);
                })
                
            }
            
        })
        .catch((err)=>{
            console.log(err);
        })
    };

    const handleSnapshot = (vmID) => {
        // Add logic for creating a snapshot of the VM with the given ID
        console.log(`Snapshot VM: ${vmID}`);
        instance.post("/transaction", {
            "dest": "/action/snapshot/"+vmID,
            "method": "POST",
            "data": ""
        })
        .then((response)=>{
            console.log(response);
            alert("Snapshot created!");
        })
        .catch((err)=>{
            console.log(err);
            alert("Failed to create snapshot");
        })
    }

    const handleReboot = (vmID) => {
        // Add logic for rebooting the VM with the given ID
        console.log(`Reboot VM: ${vmID}`);
        instance.post("/transaction", {
            "dest": "/action/softreboot/"+vmID,
            "method": "POST",
            "data": ""
        })
        .then((response)=>{
            console.log(response);
            alert("VM rebooted!");
            
        })
        .catch((err)=>{
            console.log(err);
            alert("Failed to reboot VM");
        })
    }

    const handlePowerOff = (vmID) => {
        // Add logic for powering off the VM with the given ID
        console.log(`PowerOff VM: ${vmID}`);
        instance.post("/transaction", {
            "dest": "/action/poweroff/"+vmID,
            "method": "POST",
            "data": ""
        })
        .then((response)=>{
            console.log(response);
            alert("VM powered off!");
        })
        .catch((err)=>{
            console.log(err);
            alert("Failed to power off VM");
        })
    }

    const handlePowerOn = (vmID) => {
        // Add logic for powering on the VM with the given ID
        console.log(`PowerOn VM: ${vmID}`);
        instance.post("/transaction", {
            "dest": "/action/poweron/"+vmID,
            "method": "POST",
            "data": ""
        })
        .then((response)=>{
            console.log(response);
            alert("VM powered on!");
        })
        .catch((err)=>{
            console.log(err);
            alert("Failed to power on VM");
        })
    }

    const handleConsole = (vmID) => {
        // Add logic for opening the console of the VM with the given ID
        console.log(`Console VM: ${vmID}`);
        instance.post("/transaction", {
            "dest": "/action/vnc/"+vmID,
            "method": "GET",
            "data": ""
        })
        .then((response)=>{
            console.log(response.data.data);
            alert("Opening console!");
            window.open(response.data.data.replace("controller","117.16.137.241"));
        })
        .catch((err)=>{
            console.log(err);
            alert("Failed to open console");
        })
    }

    return (
        <div className="vm-list">
            
            <ul>
                {data.resources.map((vm, index) => (
                    <li key={index} className="vm-item">
                        <strong>Name:</strong> {vm.name}
                        <strong>IP:</strong> {vm.instances[0].attributes.access_ip_v4}
                        <strong>Flavor:</strong> {vm.instances[0].attributes.flavor_name}
                        <strong>Status:</strong> {vm.instances[0].attributes.power_state}
                        <strong>KeyPair:</strong> {vm.instances[0].attributes.key_pair}
                        {/* Delete and Manage buttons */}
                        <div className="vm-buttons">
                            <button className="delete-button" onClick={() => handleDelete(vm.instances[0].attributes.id)}>Delete</button>
                            <button className="snapshot-button" onClick={() => handleSnapshot(vm.instances[0].attributes.id)}>Snapshot</button>
                            <button className="reboot-button" onClick={() => handleReboot(vm.instances[0].attributes.id)}>Reboot</button>
                            <button className="poweroff-button" onClick={() => handlePowerOff(vm.instances[0].attributes.id)}>PowerOff</button>
                            <button className="poweron-button" onClick={() => handlePowerOn(vm.instances[0].attributes.id)}>PowerOn</button>
                            <button className="console-button" onClick={() => handleConsole(vm.instances[0].attributes.id)}>Console</button>
                        </div>
                        <br />
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default VmList;