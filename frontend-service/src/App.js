import React from "react";
import PingTable from "./PingTable";
import 'antd/dist/reset.css';

const App = () => {
  return (
    <div style={{ padding: '20px' }}>
      <h1>Ping Results</h1>
      <PingTable />
    </div>
  );
};

export default App;
