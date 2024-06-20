import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'react-bootstrap';
import './index.css'; // CSS 파일을 통해 스타일 적용

const ShowMonitoring = () => {
  return (
    <div className="monitoring">    
      <div className="monitoring-webview">
        <a href ="http://117.16.137.217:5601/app/dashboards#/view/34e8b060-2d90-11ef-999c-590a08d8265c?embed=true&_g=(refreshInterval:(pause:!t,value:60000),time:(from:now-15m,to:now))&_a=()" target="_blank">
          <Button variant="primary">View Full Screen</Button>
        </a>
        <iframe 
          src="http://117.16.137.217:5601/app/home#/" 
          title="Web Monitoring"
          width="100%" 
          height="600px" 
          style={{ border: 'none' }}
        >
        </iframe>

      </div>
  
    </div>
  );
};

export default ShowMonitoring;
