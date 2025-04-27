import React, { useState } from 'react';
import axios from 'axios';

function Upload() {
  const [file, setFile] = useState(null);
  const [jsonText, setJsonText] = useState('');

  const uploadFile = async () => {
    if (!file) {
      alert("Please select a file first!");
      return;
    }
    const formData = new FormData();
    formData.append("file", file);

    try {
      await axios.post('http://localhost:8080/upload', formData);
      alert("File uploaded successfully!");
    } catch (err) {
      alert("File upload failed!");
    }
  };

  const uploadJSON = async () => {
    if (!jsonText.trim()) {
      alert("Please enter valid JSON");
      return;
    }

    try {
      const parsed = JSON.parse(jsonText);
      await axios.post('http://localhost:8080/upload', parsed, {
        headers: { 'Content-Type': 'application/json' }
      });
      alert("JSON uploaded successfully!");
    } catch (err) {
      alert("Invalid JSON or upload failed!");
    }
  };

  return (
    <div className="container">
      <h2>Upload Parquet File</h2>
      <div className="upload-bar">
        <input type="file" onChange={(e) => setFile(e.target.files[0])} />
        <button onClick={uploadFile}>Upload File</button>
      </div>

      <h2>Or Upload JSON</h2>
      <textarea
        rows={8}
        cols={60}
        placeholder="Paste JSON here..."
        value={jsonText}
        onChange={(e) => setJsonText(e.target.value)}
        style={{ padding: '10px', borderRadius: '6px', marginTop: '10px' }}
      />
      <br />
      <button onClick={uploadJSON} style={{ marginTop: '10px' }}>Upload JSON</button>
    </div>
  );
}

export default Upload;
