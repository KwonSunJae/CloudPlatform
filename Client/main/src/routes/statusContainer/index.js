import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from '../../components/navigationbar';
import Footer from "../../components/footer";
import ContainerStatus from "../../components/container/status";

export default function ContainerStat() {
  
  return (
    
    <div>
        <NavigationBar/>
        <ContainerStatus/>
        <Footer/>
    </div>
    
  );
}