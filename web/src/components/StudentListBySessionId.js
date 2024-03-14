import React, { useState, useEffect } from 'react';

const StudentListBySessionId = () => {
  const [studentList, setStudentList] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [classList, setClassList] = useState([]);
  const [selectedSessionId, setSelectedSessionId] = useState('');

  useEffect(() => {
    const fetchSessionIds = async () => {
      try {
        const response = await fetch('http://127.0.0.1:5000/student/clslist');
        if (response.ok) {
          const data = await response.json();
          setClassList(data.data);
        } else {
          console.error('Failed to fetch session IDs');
        }
      } catch (error) {
        console.error('Error fetching session IDs:', error);
      }
    };

    fetchSessionIds();
  }, []);

  const handleSubmit = async () => {
    try {
      setLoading(true);

      const response = await fetch('http://127.0.0.1:5000/student/stdlistbysessionid', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ sessionId: selectedSessionId })
      });

      if (!response.ok) {
        throw new Error('Failed to fetch student list');
      }

      const data = await response.json();
      setStudentList(data.message);
      setLoading(false);
    } catch (error) {
      setError(error.message);
      setLoading(false);
    }
  };

  const handleSessionChange = (event) => {
    setSelectedSessionId(event.target.value);
  };
  // Function to format Unix timestamp to "Wednesday, 12 March 2024" format
  const formatDate = (dob) => {
    if (!dob) return ''; // Return empty string if dob is not provided
    const date = new Date(dob.seconds * 1000); // Convert seconds to milliseconds

    // Array of month names
    const months = [
      'January', 'February', 'March', 'April', 'May', 'June', 'July',
      'August', 'September', 'October', 'November', 'December'
    ];

    // Array of day names
    const days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];

    // Get day of the week, day of the month, and month name
    const dayOfWeek = days[date.getDay()];
    const dayOfMonth = date.getDate();
    const monthName = months[date.getMonth()];
    const year = date.getFullYear();

    // Concatenate and return the formatted date string
    return `${dayOfWeek}, ${dayOfMonth} ${monthName} ${year}`;
  };

  // Function to calculate age based on date of birth
const calculateAge = (dob) => {
  if (!dob) return ''; // Return empty string if dob is not provided
  const date = new Date(dob.seconds * 1000); // Convert seconds to milliseconds
  const today = new Date(); // Get current date

  // Calculate age based on the difference in years
  let age = today.getFullYear() - date.getFullYear();

  // Adjust age if the birthday has not occurred yet this year
  const isBirthdayPast = today.getMonth() > date.getMonth() || 
    (today.getMonth() === date.getMonth() && today.getDate() >= date.getDate());
  if (!isBirthdayPast) {
    age--;
  }

  return age;
};



  return (
    <div>
      <h2>Fetch Student List By Session ID</h2>
      <div>
        <label>Select Session ID:</label>
        <select value={selectedSessionId} onChange={handleSessionChange}>
          <option value="">-- Select Session --</option>
          {classList.map(session => (
            <option key={session.sessionId} value={session.sessionId}> {`${session.class} - ${session.year}`}</option>
          ))}
        </select>
        <button onClick={handleSubmit} disabled={!selectedSessionId || loading}>Fetch Student List</button>
      </div>
      {loading && <p>Loading...</p>}
      {error && <p>Error: {error}</p>}
      {studentList !== null && studentList.length > 0 && (
        <div className="table-container">
          <table className="student-table" style={{ borderCollapse: 'collapse', width: '100%', margin: '20px 0', border: '1px solid #ddd' }}>
            <thead>
              <tr style={{ backgroundColor: '#f2f2f2' }}>
                {/* <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>ID</th> */}
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>No</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Name</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>ID</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Name BN</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Father's Name</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Mothers's Name</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Gender</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Mobile</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Address</th>
                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>DOB</th>
                {/* Add more table headers for other student details */}
              </tr>
            </thead>
            <tbody>
              {studentList != null && studentList.map((student, index) => (
                <tr key={student.id}>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{index + 1}</td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>
                     <a href={`student/${student.id}/profile`}>{student.name}</a>
                    </td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.std_id}</td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.name_bn}</td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.fathers_name}</td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.mothers_name}</td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.gender}</td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.mobile_number}</td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.address}</td>
                  <td style={{ padding: '12px', border: '1px solid #ddd' }}>{formatDate(student.dob)} - ({calculateAge(student.dob)})</td>
                  {/* Add more table cells for other student details */}
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
};

export default StudentListBySessionId;
