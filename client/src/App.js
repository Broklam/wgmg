import {Route, Routes} from "react-router-dom"
import { Home } from "./pages/Home";
import {Login} from "./pages/Login"
import {About} from "./pages/About"
import React  from 'react';
import Layout from "./layout/layout";


function App() {
  return (
    <Layout>
    <Routes> 

      <Route path="/" element={<Home/>}/>
      <Route path="/about" element={<About/>}/>
      <Route path="/login" element={<Login/>}/>
      <Route/>
    </Routes>
    </Layout>
  );
}

export default App;
