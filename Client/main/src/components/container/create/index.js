// ContainerCreatePage.js
import React, { useState } from 'react';
import DeploymentForm from '../deployment/index';
import ServiceForm from '../service/index';
import "./index.css";
const ContainerCreatePage = () => {
  const [selectedOption, setSelectedOption] = useState('deployment');

  const handleOptionChange = (option) => {
    setSelectedOption(option);
  };

  return (
    <div className="container-create-page">
      <h2>Create Container</h2>
      <div>
        <label>
          <input
            type="radio"
            value="deployment"
            checked={selectedOption === 'deployment'}
            onChange={() => handleOptionChange('deployment')}
          />
          Deployment
        </label>
        <label>
          <input
            type="radio"
            value="service"
            checked={selectedOption === 'service'}
            onChange={() => handleOptionChange('service')}
          />
          Service
        </label>
      </div>

      {selectedOption === 'deployment' ? <DeploymentForm /> : <ServiceForm />}
    </div>
  );
};

export default ContainerCreatePage;
