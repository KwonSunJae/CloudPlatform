// VmList.js
import React from 'react';
import instance from '../../../apis/index'

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

    const handleManage = (vmName) => {
        // Add logic for managing the VM with the given name
        console.log(`Manage VM: ${vmName}`);
    };

    return (
        <div className="vm-list">
            
            <ul>
                {data.resources.map((vm, index) => (
                    <li key={index} className="vm-item">
                        <strong>Name:</strong> {vm.name}<br />
                        <strong>IP:</strong> {vm.instances[0].attributes.access_ip_v4}<br />
                        <strong>Flavor:</strong> {vm.instances[0].attributes.flavor_name}<br />
                        <strong>Status:</strong> {vm.instances[0].attributes.power_state}<br />
                        <strong>KeyPair:</strong> {vm.instances[0].attributes.key_pair} <br />
                        {/* Delete and Manage buttons */}
                        <div className="vm-buttons">
                            <button className="delete-button" onClick={() => handleDelete(vm.name)}>Delete</button>
                            <button className="manage-button" onClick={() => handleManage(vm.name)}>Manage</button>
                        </div>
                        <br />
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default VmList;