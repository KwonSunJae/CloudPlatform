// ServiceForm.js
import React,{useState,useEffect} from 'react';
import instance from '../../../../apis/instance';
const ServiceForm = () => {
  const [serviceName,setServiceName] = useState('');
  const [deploymentName, setDeploymentName] = useState('');
  const [protocol, setProtocol] = useState('');
  const [targetPort, setTargetPort] = useState('');
  const [port, setPort] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    
    const datas = JSON.stringify(
      {
        apiVersion: "v1" ,
        kind: "Service",
        metadataName: serviceName,
        specClusterIP: "none",
        specExternalname: "none",
        specPortsNodeport: targetPort,
        specPortsPort: port,
        specPortsProtocol: protocol ,
        specPortsTargetport: targetPort ,
        specSelectorApp: "none",
        specSelectorType: "none",
        specType: "NodePort"
      }
    )
    instance.post("/transaction",{
      dest: "/service",
      method: "POST",
      data: datas
    }
    )
      .then(response => {
        console.log(response.data);
        alert("생성완료");
        
      })
      .catch(error => {
        console.error(error);
      });


    // Implement logic to create the virtual machine
  };
  return (
    <div className="deployment-form">
      
      <h2>Create New Service</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Service Name:
          <input
            type="text"
            value={serviceName}
            onChange={(e) => setServiceName(e.target.value)}
          />
        </label>

        <label>
          Deployment Name:
          <input
            type="text"
            value={deploymentName}
            onChange={(e) => setDeploymentName(e.target.value)}
          />
        </label>

        <label>
          Port:
          <input
            type="text"
            value={port}
            onChange={(e) => setPort(e.target.value)}
          />
        </label>

        <label>
          Protocol:
          <input
            type="text"
            checked={protocol}
            onChange={(e) => setProtocol(e.target.value)}
          />
        </label>

        <label>
          TargetPort :
          <input
            type="text"
            checked={targetPort}
            onChange={(e) => setTargetPort(e.target.value)}
          />
        </label>


        <button type="submit">Create Service</button>
      </form>
    </div>
  );
};

export default ServiceForm;
