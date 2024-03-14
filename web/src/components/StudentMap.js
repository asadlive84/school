import React from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';

const StudentMap = ({students}) => {
  const panchoKoron = [22.50872, 89.79454]
  return (
    <MapContainer center={panchoKoron} zoom={13} style={{ height: '400px' }}>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      {/* Add markers for each student */}
      {students.map((student, index) => (
        <Marker key={index} position={panchoKoron}>
          <Popup>{student.name}</Popup>
          <Popup>{student.name_bn}</Popup>
        </Marker>
      ))}
    </MapContainer>
  );
};

export default StudentMap;
