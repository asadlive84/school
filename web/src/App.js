import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import FileUploadForm from './components/studentUpload';
import Home from './components/Home';
import StudentList from './components/StudentList';
import StudentListBySessionId from './components/StudentListBySessionId';
import StudentProfile from './components/StudentProfile';
import Layout from './components/Layout';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<Layout />}>
          <Route path="/home" element={<Home />} />
          <Route path="/student/:uuid/profile" element={ <StudentProfile/> } />
          <Route path="/students" element={<StudentList />} />
          <Route path="/stdlistsessionid" element={<StudentListBySessionId />} />
          <Route path="/upload" element={<FileUploadForm />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
