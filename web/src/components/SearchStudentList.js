import React from 'react';

function StudentList({ students }) {
  return (
    <div>
      <h2>Student List</h2>
      {students && students.length > 0 ? (
        <table className="student-table">
          <thead>
            <tr>
              <th>Name</th>
              <th>ID</th>
              <th>Name (Bangla)</th>
              <th>Father's Name</th>
              <th>Mother's Name</th>
              <th>Date of Birth</th>
              <th>Gender</th>
              <th>Blood Group</th>
              <th>Mobile Number</th>
              <th>Session</th>
              <th>Class Name</th>
              <th>Village/Road</th>
              <th>Union</th>
              <th>Upazilla</th>
              <th>District</th>
            </tr>
          </thead>
          <tbody>
            {students.map(student => (
              <tr key={student.id}>
                <td>
                  <a href={`student/${student.id}/profile`}>{student.name}</a>
                </td>
                <td>{student.std_id}</td>
                <td>{student.name_bn}</td>
                <td>{student.fathers_name}</td>
                <td>{student.mothers_name}</td>
                <td>{new Date(student.dob.seconds * 1000).toLocaleDateString()}</td>
                <td>{student.gender}</td>
                <td>{student.blood_group}</td>
                <td>{student.mobile_number}</td>
                <td>{student.session}</td>
                <td>{student.class_name}</td>
                <td>{student.village_or_road}</td>
                <td>{student.unionName}</td>
                <td>{student.upzilla}</td>
                <td>{student.district}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>No students available</p>
      )}
      <style jsx>{`
        .student-table {
          width: 100%;
          border-collapse: collapse;
        }
        .student-table th, .student-table td {
          border: 1px solid #ddd;
          padding: 8px;
          text-align: left;
        }
        .student-table th {
          background-color: #f2f2f2;
        }
        .student-table tr:nth-child(even) {
          background-color: #f2f2f2;
        }
      `}</style>
    </div>
  );
}

export default StudentList;
