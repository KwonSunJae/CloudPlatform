import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../../components/navigationbar";
import Footer from "../../../components/footer";
import AdminVM from "../../../components/admin/vm";
import AdminNavbar from "../../../components/admin/navbar";



export default function AdminPageVM() {
  
  return (
    
    <div>
    <NavigationBar />
    <AdminNavbar />
    <AdminVM />
    <Footer />
    </div>
    
  );
}