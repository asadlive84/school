import React, { useState, useEffect } from 'react';

const FileUploadForm = () => {
  const [selectedFile, setSelectedFile] = useState(null);
  const [selectedClass, setSelectedClass] = useState('');
  const [classList, setClassList] = useState([]);

  useEffect(() => {
    const fetchClassList = async () => {
      try {
        const response = await fetch('http://127.0.0.1:5000/student/clslist');
        if (response.ok) {
          const data = await response.json();
          setClassList(data.data);
        } else {
          console.error('Failed to fetch class list');
        }
      } catch (error) {
        console.error('Error fetching class list:', error);
      }
    };

    fetchClassList();
  }, []);

  const handleFileChange = (event) => {
    setSelectedFile(event.target.files[0]);
  };

  const handleClassChange = (event) => {
    const value = event.target.value;
    console.log("Selected class:", value);
    setSelectedClass(value);
};


  const handleUpload = async () => {
    if (!selectedFile) {
      alert('Please select a file');
      return;
    }

    if (!selectedClass) {
      alert('Please select a class');
      return;
    }

    const formData = new FormData();
    formData.append('file', selectedFile);
    formData.append('sessionYear', selectedClass);

    try {
      const response = await fetch('http://localhost:5000/student/uploadfile', {
        method: 'POST',
        body: formData,
      });

      if (response.ok) {
        alert('File uploaded successfully');
      } else {
        alert('Failed to upload file');
      }
    } catch (error) {
      console.error('Error uploading file:', error);
      alert('An error occurred while uploading the file');
    }
  };

  return (
    <div className="upload-container">
      <h2>Upload CSV File</h2>
      <input type="file" accept=".csv" onChange={handleFileChange} />
      <div>
        <label>Class:</label>
        <select name="classSession" value={selectedClass} onChange={handleClassChange}>
          {classList !== null && classList && classList.map(cls => (
            <option key={cls.sessionId} value={cls.sessionId}>
              {`${cls.class}-${cls.year}`}
            </option>
          ))}
          {/* Add another option */}
          <option value="other">Other</option>
        </select>
      </div>
      <button onClick={handleUpload}>Upload</button>
    </div>
  );
  
};

export default FileUploadForm;
