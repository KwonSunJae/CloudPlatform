// DeploymentForm.js
import React, { useState } from 'react';
import instance from '../../../../apis';

const DeploymentForm = () => {
  const [deploymentName, setDeploymentName] = useState('');
  const [hostName, setHostName] = useState('');
  const [subdomain, setSubdomain] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [imagePullPolicy, setImagePullPolicy] = useState('');
  const [port, setPort] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    instance.post("/deployment", {
      ApiVersion : "apps/v1",
      Kind : "Deployment",
      Metadata_name : deploymentName,
      Metadata_labels_app: deploymentName,
      Spec_selector_matchLabels_app  : deploymentName,
      Spec_template_metadata_labels_app : deploymentName,
      Spec_template_spec_hostname: hostName,
      Spec_template_spec_subdomain: subdomain,
      Spec_template_spec_containers_image : imageUrl,
      Spec_template_spec_containers_imagePullPolicy: imagePullPolicy,
      Spec_template_spec_containers_name:deploymentName,
      Spec_template_spec_containers_ports_containerPort:port
    })
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
          Image Pull Policy:
          <input
            type="text"
            checked={imagePullPolicy}
            onChange={(e) => setImagePullPolicy(e.target.value)}
          />
        </label>

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
