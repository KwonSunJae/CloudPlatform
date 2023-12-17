import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from '../../components/navigationbar';
import Footer from "../../components/footer";
import VmStatus from "../../components/vm/status";

export default function VMStat() {
  
  return (
    
    <div>
        <NavigationBar/>
        <VmStatus/>
        <Footer/>
    </div>
    
  );
}