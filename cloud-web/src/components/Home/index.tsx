import * as React from 'react';
import Accordion from '@mui/material/Accordion';
import AccordionSummary from '@mui/material/AccordionSummary';
import AccordionDetails from '@mui/material/AccordionDetails';
import Typography from '@mui/material/Typography';
import Infrastructure from './Infrastructure';
import generateRandomHexColor from '../../utils/randomColor';
import { useEffect } from 'react';
import { API_URL } from "../../constants";
const Home: React.FC = () => {
  const [infrastructureData, setInfrastructureData] = React.useState([]);
  useEffect(() => {
    const fetchData = async () => {
      if (window.isFocused) {
        try {
          fetch(`${API_URL}/infrastructures`)
            .then(response => response.json())
            .then(data => {
              console.log(data.infrastructures);
              if (infrastructureData.length < data.infrastructures.length) {
              setInfrastructureData(data.infrastructures);
            }
          });
      } catch (err) {
        console.error("Error fetching infrastructure data:", err);
      }}
    };

    fetchData();
    const intervalId = setInterval(fetchData, 5000);

    return () => clearInterval(intervalId);
  }, []);

  return (
    <div>
      {infrastructureData.length && infrastructureData.map((infra: any, index) => (
        <Infrastructure key={index} color={generateRandomHexColor(infra.Type)} header={infra.Type + " - " + infra.Name} dataObject={infra}>
          {/* <Typography>{infra?.Details}</Typography> */}
        </Infrastructure>
      ))}
    </div>
  );
};

export default Home;
