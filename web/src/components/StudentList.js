import React, { useState, useEffect } from 'react';
import MapComponent from './MapComponent';
import StudentMap from './StudentMap';

function StudentList() {
    const [students, setStudents] = useState([]);
    const [error, setError] = useState(null);
    const [filter, setFilter] = useState({ name: '', id: '', className: '' });

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

    const filteredStudents = students.filter(student =>
        student.name.toLowerCase().includes(filter.name.toLowerCase()) &&
        student.id.toLowerCase().includes(filter.id.toLowerCase()) &&
        student.class_name.toLowerCase().includes(filter.className.toLowerCase())
    );

    const handleFilterChange = (event) => {
        const { name, value } = event.target;
        setFilter(prevFilter => ({ ...prevFilter, [name]: value }));
    };

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div>
            <div>
                {/* {students.length > 0 && <MapComponent studentData={students} />} */}
            </div>

            <div>
                <h2>Filter Students</h2>
                <div>
                    <input
                        type="text"
                        name="name"
                        placeholder="Search by name"
                        value={filter.name}
                        onChange={handleFilterChange}
                    />
                </div>
                <div>
                    <input
                        type="text"
                        name="id"
                        placeholder="Search by ID"
                        value={filter.id}
                        onChange={handleFilterChange}
                    />
                </div>
                <div>
                    <input
                        type="text"
                        name="className"
                        placeholder="Search by class name"
                        value={filter.className}
                        onChange={handleFilterChange}
                    />
                </div>
            </div>

            {students === null ? (
                <p>Student list is empty</p>
            ) : (
                <div>
                    <h1>Student List</h1>
                    <table style={{ borderCollapse: 'collapse', width: '100%', margin: '20px 0', border: '1px solid #ddd' }}>
                        <thead>
                            <tr style={{ backgroundColor: '#f2f2f2' }}>
                                <th style={{ padding: '12px', textAlign: 'left', border: '1px solid #ddd' }}>#</th>
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
                            {filteredStudents.map((student, index) => (
                                <tr key={student.id}>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>{index + 1}</td>
                                    <td style={{ padding: '12px', border: '1px solid #ddd' }}>
                                       <a href={`student/${student.id}/profile`}>{student.name}</a>
                                    </td>
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
