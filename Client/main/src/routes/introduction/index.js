import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../components/navigationbar";
import Footer from "../../components/footer";

import ShowIntroduction from "../../components/introduction";



export default function Introduction() {
  
  return (
    
    <div>
    <NavigationBar />
    <ShowIntroduction />
    <Footer />
    </div>
    
  );
}