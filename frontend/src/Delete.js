import React, { useState } from 'react';
import axios from 'axios';

function DeleteEvent() {
  const [eventId, setEventId] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const [showTooltip, setShowTooltip] = useState(false);

  // Delete event handler
  const deleteEvent = async () => {
    if (!eventId.trim()) {
      alert("Please enter a valid event ID!");
      return;
    }

    setLoading(true); // Show loading state
    setMessage('');  // Clear previous messages
    setShowTooltip(false); // Hide tooltip initially

    try {
      // Send DELETE request to backend to delete event
      const res = await axios.delete(`http://localhost:8080/delete?event_id=${eventId}`);

      // On success, show the success message and clear loading state
      setLoading(false);
      setMessage(res.data); // Use the server response message
      setShowTooltip(true); // Show tooltip

      // Hide tooltip after 5 seconds
      setTimeout(() => setShowTooltip(false), 5000);
    } catch (error) {
      // Handle error during deletion
      setLoading(false);
      setMessage("Error deleting event. Please try again.");
      setShowTooltip(true); // Show tooltip

      // Hide tooltip after 5 seconds
      setTimeout(() => setShowTooltip(false), 5000);
    }
  };

  return (
    
    <div className="container">
      <h2>Delete Event</h2>

      {/* Event ID input */}
      <input
        type="text"
        value={eventId}
        onChange={(e) => setEventId(e.target.value)}
        placeholder="Enter Event ID"
        style={{
          padding: '8px',
          borderRadius: '4px',
          border: '1px solid #ccc',
          width: '300px',
        }}
      />

      {/* Delete button */}
      <button
        onClick={deleteEvent}
        style={{
          marginLeft: 10,
          padding: '8px 16px',
          borderRadius: '4px',
          cursor: 'pointer',
          backgroundColor: '#3498db',
          color: 'white',
        }}
      >
        {loading ? 'Deleting...' : 'Delete'}
      </button>

      {/* Tooltip for API response */}
      {showTooltip && (
        <div 
          style={{
            position: 'absolute',
            top: '40px',
            left: '0',
            backgroundColor: '#333',
            color: 'white',
            padding: '8px 12px',
            borderRadius: '4px',
            zIndex: 999,
            fontSize: '14px',
            opacity: 1,
            transition: 'opacity 0.3s ease',
          }}
        >
          {message}
        </div>
      )}
    </div>
  );
}

export default DeleteEvent;
