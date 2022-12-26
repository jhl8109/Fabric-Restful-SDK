import './App.css';
import React from 'react';
import {BrowserRouter, Route} from "react-router-dom";
import Channel from './components/channel'
import Chaincode from './components/chaincode';

function App() {
  return (
    <div className="App">
       <BrowserRouter>
        <Route path = '/channels' component = {Channel}/> 
        <Route path = '/chaincodes' component = {Chaincode}/> 
       </BrowserRouter>
      
    </div>
  );
}

export default App;
