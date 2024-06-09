import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'react-bootstrap';
import './index.css'; // CSS 파일을 통해 스타일 적용

const ShowMonitoring = () => {
  return (
    <div className="monitoring">
      <div className="monitoring-header">
        <h1>Grafana Monitoring</h1>
      </div>
      
      <div className="monitoring-webview">
        <iframe 
          src="http://117.16.136.172:3000/swagger/index.html" 
          title="Web Monitoring"
          width="100%" 
          height="600px" 
          style={{ border: 'none' }}
        ></iframe>
      </div>
  
    </div>
  );
};

export default ShowMonitoring;
