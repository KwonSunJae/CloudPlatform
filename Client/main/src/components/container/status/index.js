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
        <div className="container">
            <h1>컨테이너 목록</h1>
            <button className="add-ct-button" onClick={handleAddCtClick}>
                컨테이너 추가
            </button>
            <div className="split-layout">
                <h2>Deployment 목록</h2>
                <div className="deployment-section">
                    {deploymentData ? (
                        <DeploymentList data={deploymentData} />
                    ) : (
                        <h1>Server error</h1>
                    )}
                </div>
                <h2>Service 목록</h2>
                <div className="service-section">
                    {serviceData ? (
                        <ServiceList data={serviceData} />
                    ) : (
                        <h1>Server Error</h1>
                    )}
                </div>
            </div>
        </div>
    );
};

export default ContainerStatus;
