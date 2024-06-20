// ContainerCreatePage.js
import React, { useState } from 'react';
import DeploymentForm from './deployment/index';
import ServiceForm from './service/index';
import "./index.css";

const ContainerCreatePage = () => {
  const [selectedOption, setSelectedOption] = useState('deployment');

  const handleOptionChange = (option) => {
    setSelectedOption(option);
  };

  return (
    <div className="container-create-page">
      <h2>Create Container</h2>
      <div className="option-buttons">
        <button
          className={selectedOption === 'deployment' ? 'selected' : ''}
          onClick={() => handleOptionChange('deployment')}
        >
          Deployment
        </button>
        <button
          className={selectedOption === 'service' ? 'selected' : ''}
          onClick={() => handleOptionChange('service')}
        >
          Service
        </button>
      </div>

      {selectedOption === 'deployment' ? <DeploymentForm /> : <ServiceForm />}
    </div>
  );
};

export default ContainerCreatePage;
