import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';


const StudentProfile = () => {



    const { uuid } = useParams(); // Extracting UUID from URL params
    const [studentProfile, setStudentProfile] = useState(null);
    console.log("==================", uuid)

    useEffect(() => {
        fetch(`http://127.0.0.1:5000/student/profile/${uuid}`)
            .then(response => response.json())
            .then(data => {
                console.log("Received data:", data.message); // Log the received data
                setStudentProfile(data.message);
            })
            .catch(error => {
                console.error('Error fetching student profile:', error);
            });
    }, [uuid]);

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

    return (
        <div>
            <h1>Student Profile</h1>
            {studentProfile !== null ? (
                <div>
                    {/* <p><strong>ID:</strong> {studentProfile.id}</p> */}
                    <p><strong>Student ID:</strong> {studentProfile.id === " " ? (<div>
                        no id
                    </div>):studentProfile.std_id}</p>
                    <p><strong>Name:</strong> {studentProfile.name}</p>
                    <p><strong>Name (Bangla):</strong> {studentProfile.name_bn}</p>
                    <p><strong>Father's Name:</strong> {studentProfile.fathers_name}</p>
                    <p><strong>Mother's Name:</strong> {studentProfile.mothers_name}</p>
                    {/* <p><strong>Date of Birth:</strong> {studentProfile.dob}</p> */}
                    <p><strong>Date of Birth:</strong> {formatDate(studentProfile.dob)}</p>
                    <p><strong>Gender:</strong> {studentProfile.gender}</p>
                    <p><strong>Blood Group:</strong> {studentProfile.blood_group}</p>
                    <p><strong>Mobile Number:</strong> {studentProfile.mobile_number}</p>
                    <p><strong>Session:</strong> {studentProfile.session}</p>
                    <p><strong>Class Name:</strong> {studentProfile.class_name}</p>
                    <p><strong>Religion:</strong> {studentProfile.religion}</p>
                    <p><strong>District:</strong> {studentProfile.district}</p>
                    <p><strong>Upzilla:</strong> {studentProfile.upzilla}</p>
                    <p><strong>Union/City:</strong> {studentProfile.unionName}</p>
                    <p><strong>Village/Road:</strong> {studentProfile.village_or_road}</p>
                </div>
            ) : 'Not loading'}
        </div>
    );

}

export default StudentProfile;
