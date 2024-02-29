import React, { useState } from 'react';

const App = () => {
  const [formData, setFormData] = useState({
    name: '',
    name_bn: '',
    std_id: '',
    fathers_name: '',
    mothers_name: '',
    session: '',
    gender: '',
    mobile_number: '',
    class_name: ''
  });

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://127.0.0.1:3000/student/insert', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
      });
      if (response.ok) {
        console.log('Data successfully submitted');
        // Reset form data if needed
        setFormData({
          name: '',
          name_bn: '',
          std_id: '',
          fathers_name: '',
          mothers_name: '',
          session: '',
          gender: '',
          mobile_number: '',
          class_name: ''
        });
      } else {
        console.error('Failed to submit data');
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  return (
    <div>
      <h1>Student Form</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">Name:</label>
          <input type="text" id="name" name="name" value={formData.name} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="name_bn">Name (Bangla):</label>
          <input type="text" id="name_bn" name="name_bn" value={formData.name_bn} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="std_id">Student ID:</label>
          <input type="text" id="std_id" name="std_id" value={formData.std_id} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="fathers_name">Father's Name:</label>
          <input type="text" id="fathers_name" name="fathers_name" value={formData.fathers_name} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="mothers_name">Mother's Name:</label>
          <input type="text" id="mothers_name" name="mothers_name" value={formData.mothers_name} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="session">Session:</label>
          <input type="text" id="session" name="session" value={formData.session} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="gender">Gender:</label>
          <input type="text" id="gender" name="gender" value={formData.gender} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="mobile_number">Mobile Number:</label>
          <input type="text" id="mobile_number" name="mobile_number" value={formData.mobile_number} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="class_name">Class Name:</label>
          <input type="text" id="class_name" name="class_name" value={formData.class_name} onChange={handleChange} />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default App;
