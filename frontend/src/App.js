import React from 'react';
import Upload from './Upload';
import Search from './Search';
import Delete from './Delete';
import './App.css';

function App() {
  return (
    <div style={{ padding: 20 }}>
      <h1>Search Engine</h1>
      <Upload />
      <Delete />
      <Search />
    </div>
  );
}

export default App;