
import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../../components/navigationbar";
import Footer from "../../../components/footer";
import AdminContainerService from "../../../components/admin/containerService";
import AdminNavbar from "../../../components/admin/navbar";



export default function AdminPageContainerService() {
  
  return (
    
    <div>
    <NavigationBar />
    <AdminNavbar />
    <AdminContainerService />
    <Footer />
    </div>
    
  );
}