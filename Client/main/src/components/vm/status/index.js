import React, { useEffect, useState } from 'react';

import VmList from './vmlist'; // Make sure to provide the correct path
import instance from '../../../apis/instance';
import SpinnerOurs from '../../sppinner';
import './index.css';

const VmStatus = () => {
  const [responseData, setResponseData] = useState(null);
  const [loading, setLoading] = useState(true);

  const handleAddVmClick = () => {
    // Redirect to /create/vm
    window.location.href = '/create/vm';
  };

  useEffect(() => {
    instance.post("/transaction", {
      "dest": "/vmstat",
      "method": "GET",
      "data": ""
    }).then((response) => {
      const datas = JSON.parse(response.data.data);
      setLoading(false);
      setResponseData(datas);
      console.log(datas);
    }).catch((error) => {
      setResponseData(null);
      setLoading(false);
      console.error(error);
    });
  }, []); // Empty dependency array to run the effect only once on mount
  if (loading) {
    return <SpinnerOurs />
  }
  return (
    <div class="container">
      <h1> VM 목록</h1>
      <button class="add-vm-button" onClick={handleAddVmClick}>VM 추가</button>
      <div class="vm-list-container">
        {responseData ? <VmList data={responseData} /> : <p>VM 목록이 없습니다.</p>}
      </div>
    </div>
  );
};

export default VmStatus;
