import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../../components/navigationbar";
import Footer from "../../../components/footer";
import AdminContainerDeployment from "../../../components/admin/containerDeployment";
import AdminNavbar from "../../../components/admin/navbar";



export default function AdminPageContainerDeployment() {
  
  return (
    
    <div>
    <NavigationBar />
    <AdminNavbar />
    <AdminContainerDeployment />
    <Footer />
    </div>
    
  );
}