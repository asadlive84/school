import React, { useState, useEffect } from 'react';

function StudentList() {
    const [students, setStudents] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        // Fetch data from the endpoint
        fetch('http://127.0.0.1:5000/student/list')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to fetch data');
                }
                return response.json();
            })
            .then(data => {
                // Set the fetched data to the state
                setStudents(data.message);
            })
            .catch(error => {
                console.error('Error fetching data:', error);
                setError(error.message);
            });
    }, []);

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div>
            {students === null ? (
                <p>Student list is empty</p>
            ) : (
                <div>
                    <h1>Student List</h1>
                    <table style={{ borderCollapse: 'collapse', width: '100%', margin: '20px 0', border: '1px solid #ddd' }}>
                        <thead>
                            <tr style={{ backgroundColor: '#f2f2f2' }}>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>#</th>
                                {/* <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>ID</th> */}
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Name</th>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Name (Bangla)</th>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Gender</th>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Class</th>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Session</th>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Fathers Name</th>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Mothers Name</th>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>Mobile Number</th>
                            </tr>
                        </thead>
                        <tbody>
                            {students && students.map((student, index) => (
                                <tr key={student.id}>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{index + 1}</td>
                                    {/* <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.id}</td> */}
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.name}</td>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.name_bn}</td>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.gender}</td>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.class_name}</td>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.session}</td>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.fathers_name}</td>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.mothers_name}</td>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{student.mobile_number}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        </div>
    );
    
    
    
}

export default StudentList;
