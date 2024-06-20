// DeploymentForm.js
import React, { useState } from 'react';
import instance from '../../../../apis/instance';

const DeploymentForm = () => {
  const [deploymentName, setDeploymentName] = useState('');
  const [Replicas, setReplicas] = useState('');
  const [hostName, setHostName] = useState('');
  const [subdomain, setSubdomain] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [containerName, setcontainerName] = useState('');
  const [port, setPort] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    const datas = JSON.stringify(
      {
        apiVersion: "apps/v1", 
        kind: "Deployment",
        metadataLabelsApp: deploymentName,
        metadataName: deploymentName,
        specReplicas: Replicas,
        specSelectorMatchlabelsApp: deploymentName,
        specTemplateMetadataLabelsApp: deploymentName,
        specTemplateSpecContainersImage: imageUrl ,
        specTemplateSpecContainersName: containerName,
        specTemplateSpecContainersPortsContainerport: port
      }
      )
    instance.post("/transaction",{
      dest: "/deployment",
      method: "POST",
      data: datas}
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
      
      <h2>Create New Deployment</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Deployment Name:
          <input
            type="text"
            value={deploymentName}
            onChange={(e) => setDeploymentName(e.target.value)}
          />
        </label>

        <label>
          Replicas:
          <input
            type="text"
            value={Replicas}
            onChange={(e) => setReplicas(e.target.value)}
          />
        </label>

        <label>
          Host Name:
          <input
            type="text"
            value={hostName}
            onChange={(e) => setHostName(e.target.value)}
          />
        </label>

        <label>
          Sub Domain:
          <input
            type="text"
            value={subdomain}
            onChange={(e) => setSubdomain(e.target.value)}
          />
        </label>

        <label>
          Image URL:
          <input
            type="text"
            checked={imageUrl}
            onChange={(e) => setImageUrl(e.target.value)}
          />
        </label>

        <label>
          Container Name:
          <input
            type="text"
            checked={containerName}
            onChange={(e) => setcontainerName(e.target.value)}
          />
        </label>

        {/* <label>
          Image Pull Policy:
          <select value={containerName} onChange={(e) => setcontainerName(e.target.value)}>
            <option value="always">always</option>
            <option value="ifNotPresent">ifNotPresent</option>
            <option value="never">never</option>
          </select>
        </label> */}

        <label>
          Port :
          <input
            type="text"
            checked={port}
            onChange={(e) => setPort(e.target.value)}
          />
        </label>

        <button type="submit">Create Deployment</button>
      </form>
    </div>
  );
};

export default DeploymentForm;
