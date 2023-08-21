import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from '../../components/navigationbar';
import Footer from "../../components/footer";
import VMCreateForm from "../../components/vm";

export default function VM() {
  
  return (
    
    <div>
      <NavigationBar/>
      <VMCreateForm/>
      <Footer/>
    </div>
    
  );
}