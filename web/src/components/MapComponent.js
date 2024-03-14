import React from 'react';
import { MapContainer, TileLayer, LayersControl, Marker, Popup, LayerGroup, Circle, FeatureGroup, Rectangle } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';

const MapComponent = ({ studentData }) => {
  const center = [22.5017, 89.7873]
  const panchoKoron = [22.50872, 89.79454]
  const position = [22.5017, 89.7873]

  const zoom = 14.5
  const height = { height: '600px' }
  return (
    <MapContainer center={position} zoom={zoom} style={height}>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      <Marker position={[22.50320, 89.78305]}>
        <Popup>
          Fulhata High School
        </Popup>
      </Marker>
      <LayersControl position="topright">
        <LayersControl.Overlay checked name="Layer group with circles">
          <LayerGroup>
            {studentData.map((student, index) => (
              <Circle
                key={index}
                center={panchoKoron}
                pathOptions={{ fillColor: 'blue' }}
                radius={15} // Adjust the radius to a smaller value
              >
                <Popup>
                  Student Name: {student.name} <br />
                  Student ID: {student.name_bn}<br />
                  Student ID: {student.address}
                </Popup>
              </Circle>
            ))}
          </LayerGroup>
        </LayersControl.Overlay>

      </LayersControl>
    </MapContainer>
  );
};

export default MapComponent;
