import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../components/navigationbar";
import Footer from "../../components/footer";

import ShowSystem from "../../components/system";



export default function System() {
  
  return (
    
    <div>
    <NavigationBar />
    <ShowSystem />
    <Footer />
    </div>
    
  );
}