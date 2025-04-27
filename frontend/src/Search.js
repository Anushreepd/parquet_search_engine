import React, { useEffect, useState } from 'react';
import axios from 'axios';

function Search() {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);

  // Fetch all events on initial load
  useEffect(() => {
    const fetchAllEvents = async () => {
      try {
        const res = await axios.get('http://localhost:8080/events');
        setResults(res.data);
      } catch (error) {
        console.error("Error fetching all events:", error);
      }
    };
    fetchAllEvents();
  }, []);

  // Search handler
  const search = async () => {
    if (!query.trim()) {
      alert("Please enter a search term!");
      return;
    }
    try {
      const res = await axios.get(`http://localhost:8080/search?query=${query}`);
      setResults(res.data);
    } catch (error) {
      console.error("Search error:", error);
    }
  };

  return (
    <div className="container">
      <h2>Search Records</h2>
      <div className="search-bar">
        <input
          type="text"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          placeholder="Search for anything"
        />
        <button onClick={search}>Search</button>
      </div>

      <ul>
        {results.map((r, idx) => (
          <li key={idx}>
            <strong>Message:</strong> {r.message} <br />
            <strong>Sender:</strong> {r.sender} <br />
            <strong>Event:</strong> {r.event}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Search;
