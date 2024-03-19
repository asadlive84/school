// ParentComponent.js
import React, { useState } from 'react';
import SearchComponent from './SearchComponent';
import StudentList from './SearchStudentList';

function ParentComponent() {
  const [searchResults, setSearchResults] = useState([]);

  const handleSearch = (searchData) => {
    // Make API call with search data
    fetch('http://127.0.0.1:5000/student/search', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(searchData)
    })
      .then(response => response.json())
      .then(data => {
        // Update search results
        setSearchResults(data.message.student);
      })
      .catch(error => console.error('Error searching:', error));
  };

  return (
    <div>
      <SearchComponent onSearch={handleSearch} />
      <StudentList students={searchResults} />
    </div>
  );
}

export default ParentComponent;
