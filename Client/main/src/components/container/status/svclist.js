// itemList.js
import React from 'react';
import instance from '../../../apis/instance'

const ServiceList = ({ data }) => {
    const handleDelete = (serviceName) => {
        // Add logic for deleting the item with the given name
        console.log(`Delete service: ${serviceName}`);
        instance.get("/service")
        .then((response)=>{
            var datas = response.data.data
            var piv ="";
            console.log(datas);
            for(var i =0; i<datas.length; i++){
                if(serviceName=== datas[i].Metadata_name){
                    piv=datas[i].Id;
                    break;
                }
            }
            if(piv === ""){
                console.error("no item");
            }
            else{
                instance.delete("/service/"+piv)
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

    const handleManage = (serviceName) => {
        // Add logic for managing the item with the given name
        console.log(`Manage item: ${serviceName}`);
    };

    console.log(data.items)

    return (
        <div className="item-list">
            
            <ul>
                {data.items.map((item, index) => (
                    <li key={index} className="item-item">
                        <strong>Name:</strong> {item.metadata.name}
                        <strong>Kind:</strong>{item.kind}
                        <strong>SpecType:</strong> {item.spec.type}
                        <strong>Timestamp:</strong> {item.metadata.creationTimestamp}
                        {/* Delete and Manage buttons */}
                        <div className="item-buttons">
                            <button className="delete-button" onClick={() => handleDelete(item.metadata.name)}>Delete</button>
                            <button className="manage-button" onClick={() => handleManage(item.metadata.name)}>Manage</button>
                        </div>
                        <br />
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ServiceList;