import React, { useState, useEffect } from 'react';

function SearchComponent({ onSearch }) {
  const [districts, setDistricts] = useState([]);
  const [upazillas, setUpazillas] = useState([]);
  const [unions, setUnions] = useState([]);
  const [villageRoads, setVillageRoads] = useState([]);
  const [classList, setClassList] = useState([]);
  const [selectedDistrict, setSelectedDistrict] = useState('');
  const [selectedUpazillas, setSelectedUpazillas] = useState([]);
  const [selectedUnions, setSelectedUnions] = useState([]);
  const [selectedVillageRoads, setSelectedVillageRoads] = useState([]);
  const [selectedSessions, setSelectedSessions] = useState([]);

  useEffect(() => {
    // Fetch district list
    fetch('http://127.0.0.1:5000/student/get-district-list')
      .then(response => response.json())
      .then(data => setDistricts(data.message))
      .catch(error => console.error('Error fetching district list:', error));

    // Fetch session data
    fetch('http://127.0.0.1:5000/student/clslist')
      .then(response => response.json())
      .then(data => setClassList(data.data))
      .catch(error => console.error('Error fetching session IDs:', error));
  }, []);

  const handleDistrictChange = (districtId) => {
    setSelectedDistrict(districtId);
    setSelectedUpazillas([]);
    setSelectedUnions([]);
    setSelectedVillageRoads([]);
    // Fetch upazillas for the selected district
    fetch(`http://127.0.0.1:5000/student/get-upazilla-list/${districtId}`)
      .then(response => response.json())
      .then(data => setUpazillas(data.message))
      .catch(error => console.error('Error fetching upazillas:', error));
  };

  const handleUpazillaChange = (upazillaId) => {
    // Fetch unions for the selected upazilla
    fetch(`http://127.0.0.1:5000/student/get-union-list/${upazillaId}`)
      .then(response => response.json())
      .then(data => setUnions(data.message))
      .catch(error => console.error('Error fetching unions:', error));
    setSelectedUpazillas([...selectedUpazillas, upazillaId]);
  };

  const handleUnionChange = (unionId) => {
    // Fetch villages or roads for the selected union
    fetch(`http://127.0.0.1:5000/student/get-village-or-road-list/${unionId}`)
      .then(response => response.json())
      .then(data => setVillageRoads(data.message))
      .catch(error => console.error('Error fetching villages or roads:', error));
    setSelectedUnions([...selectedUnions, unionId]);
  };

  const handleVillageRoadChange = (villageRoadId) => {
    // Add or remove the selected village or road based on its current state
    if (selectedVillageRoads.includes(villageRoadId)) {
      setSelectedVillageRoads(selectedVillageRoads.filter(id => id !== villageRoadId));
    } else {
      setSelectedVillageRoads([...selectedVillageRoads, villageRoadId]);
    }
  };

  const handleSessionChange = (e) => {

    console.log(e.target.selectedOptions)
    const selectedOptions = Array.from(e.target.selectedOptions, option => option.value);
    setSelectedSessions(selectedOptions);
  };



  const handleSearch = () => {
    // Make API call with selected options for search
    const searchData = {
      district: selectedDistrict,
      upazillas: selectedUpazillas,
      unions: selectedUnions,
      villageRoads: selectedVillageRoads,
      sessionIds: selectedSessions
    };
    onSearch(searchData);
  };

  return (
    <div className="search-container">
      <div className="search-item">
        <label htmlFor="district">Select District:</label>
        <select id="district" value={selectedDistrict} onChange={(e) => handleDistrictChange(e.target.value)}>
          <option value="">Select District</option>
          {districts.map(district => (
            <option key={district.id} value={district.id}>{district.name}</option>
          ))}
        </select>
      </div>
      <div className="search-item">
        <label htmlFor="upazilla">Select Upazilla:</label>
        {(selectedDistrict && upazillas.length > 0) && (
          <select id="upazilla" multiple value={selectedUpazillas} onChange={(e) => handleUpazillaChange(e.target.value)}>
            {upazillas.map(upazilla => (
              <option key={upazilla.id} value={upazilla.id}>{upazilla.name}</option>
            ))}
          </select>
        )}
      </div>
      <div className="search-item">
        <label htmlFor="union">Select Union:</label>
        {(selectedDistrict && unions.length > 0) && (
          <select id="union" multiple value={selectedUnions} onChange={(e) => handleUnionChange(e.target.value)}>
            {unions.map(union => (
              <option key={union.id} value={union.id}>{union.name}</option>
            ))}
          </select>
        )}
      </div>
      <div className="search-item">
        <label htmlFor="villageRoad">Select Village/Road:</label>
        {(selectedDistrict && villageRoads.length > 0) && (
          <select id="villageRoad" multiple value={selectedVillageRoads} onChange={(e) => handleVillageRoadChange(e.target.value)}>
            {villageRoads.map(villageRoad => (
              <option key={villageRoad.id} value={villageRoad.id}>{villageRoad.name}</option>
            ))}
          </select>
        )}
      </div>



      <div className="search-item">
        <label htmlFor="session">Select Session:</label>
        <select id="session" multiple value={selectedSessions} onChange={handleSessionChange}>
          {classList.map(session => (
            <option key={session.sessionId} value={session.sessionId}>{`${session.class}-${session.year}`}</option>
          ))}
        </select>
      </div>
      <div className="search-item">
        <button onClick={handleSearch}>Search</button>
      </div>
      <style jsx>{`
        .search-container {
          display: flex;
          flex-wrap: wrap;
          gap: 10px;
        }
        .search-item {
          flex: 1 1 auto;
          display: flex;
          align-items: center;
        }
        .search-item label {
          margin-right: 10px;
        }
        .search-item select, .search-item button {
          flex: 1 1 auto;
          padding: 5px;
          border-radius: 5px;
        }
        .search-item select:focus {
          outline: none;
          border-color: #007bff;
          box-shadow: 0 0 0 0.2rem rgba(0,123,255,.25);
        }
      `}</style>
    </div>
  );
}

export default SearchComponent;
