import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../components/navigationbar";
import Footer from "../../components/footer";

import ShowMonitoring from "../../components/monitoring";



export default function Monitoring() {
  
  return (
    
    <div>
    <NavigationBar />
    <ShowMonitoring />
    <Footer />
    </div>
    
  );
}