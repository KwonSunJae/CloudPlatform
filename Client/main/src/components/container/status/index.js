// ContainerStatus.js

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
        // Redirect to /create/container
        window.location.href = '/create/container';
    };

    useEffect(() => {
        instance
            .post('/transaction', {
                dest: '/servicestat',
                method: 'GET',
                data: '',
            }
            )
            .then((response) => {
                var datas = JSON.parse(response.data.data);
                setServiceData(datas);
                console.log(serviceData);
            })
            .catch((err) => {
                console.log(err);
            });
        instance
            .post('/transaction', {
                dest: '/deploymentstat',
                method: 'GET',
                data: '',
            }

            )
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
        // <div className="contianer-status">
        <div class="container">
            <h1>컨테이너 목록</h1>
            <button className="add-ct-button" onClick={handleAddCtClick}>
                컨테이너 추가
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
