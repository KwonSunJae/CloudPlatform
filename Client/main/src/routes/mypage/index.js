import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../components/navigationbar";
import Footer from "../../components/footer";

import Mypage from "../../components/user/mypage";



export default function MyPage() {
  
  return (
    
    <div>
    <NavigationBar />
    <Mypage />
    <Footer />
    </div>
    
  );
}