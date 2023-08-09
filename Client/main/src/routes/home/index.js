import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import Main  from  '../../components/main';
import NavigationBar from '../../components/navigationbar';
import Footer from "../../components/footer";

export default function Home() {
  
  return (
    
    <div>
      <NavigationBar/>
      <Main/>
      <Footer/>
    </div>
    
  );
}