import React, { useEffect, useState} from 'react';

import VmList from './vmlist'; // Make sure to provide the correct path
import instance from '../../../apis/instance';
import './index.css';

const VmStatus = () => {
  const [responseData, setResponseData] = useState(null);
  

  const handleAddVmClick = () => {
    // Redirect to /create/vm
    window.location.href = '/create/vm';
  };

  useEffect(() => {
    instance
      .get('/vmstat')
      .then((response) => {
        var datas = JSON.parse(response.data.data);
        setResponseData(datas);
        console.log(datas.resources);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []); // Empty dependency array to run the effect only once on mount

  return (
    <div class = "container">
      <h1> VM 목록</h1>
      <button class="add-vm-button" onClick={handleAddVmClick}>VM 추가</button>
      {responseData ? (
        <VmList data={responseData} />
      ) : (
        <p>Server Internal Error</p>
        // You can replace this with a loading spinner or any other loading indicator
      )}
    </div>
  );
};

export default VmStatus;
