// VmStatus.js

import React, { useEffect, useState } from 'react';
import DeploymentList from './deplist';
import ServiceList from './svclist'; // Make sure to provide the correct path
import instance from '../../../apis/instance';
import './index.css';

const ContainerStatus = () => {
    const [responseData, setResponseData] = useState(null);
    const [deploymentData, setDeploymentData] = useState(null);
    const [serviceData, setServiceData] = useState(null);

    const handleAddCtClick = () => {
        // Redirect to /create/vm
        window.location.href = '/create/container';
    };

    useEffect(() => {
        instance
            .get('/servicestat')
            .then((response) => {
                var datas = JSON.parse(response.data.data);
                setServiceData(datas);
                console.log(serviceData);
            })
            .catch((err) => {
                console.log(err);
            });
        instance
            .get('/deploymentstat')
            .then((response) => {
                var datas = JSON.parse(response.data.data);
                console.log(datas);
                setDeploymentData(datas);
                console.log(deploymentData);
            })
            .catch((err) => {
                console.log(err);
            });


    }, []); // Empty dependency array to run the effect only once on mount

    return (
        <div className="vm-status">
            <h1>Kubernetes Status</h1>
            <button className="add-ct-button" onClick={handleAddCtClick}>
                Add Container
            </button>
            <div className="split-layout">
                {deploymentData? (<DeploymentList data={deploymentData} title="Deployment List" />):(<h1>Server error</h1>)}
                {serviceData ? (
                    <ServiceList data={serviceData} title="Service List" />
                ) : (<h1> Server Error</h1>)}
            </div>
        </div>
    );
};

export default ContainerStatus;
