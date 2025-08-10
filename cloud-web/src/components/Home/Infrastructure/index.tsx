import * as React from 'react';
import Accordion from '@mui/material/Accordion';
import AccordionSummary from '@mui/material/AccordionSummary';
import AccordionDetails from '@mui/material/AccordionDetails';
import Typography from '@mui/material/Typography';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import JSONPrinter from '../JSONPrinter';
import { PropsWithChildren } from 'react';

interface InfrastructureBaseProps {
  color?: string;
  header?: string;
  details?: string;
  dataObject: any;
}
type InfrastructureProps = PropsWithChildren<InfrastructureBaseProps>;

const Infrastructure: React.FC<InfrastructureProps> = ({ color = "", header = "", details = "", children, dataObject }) => {

    return (
        <Accordion>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel1-content"
                id="panel1-header"
                sx={{
                    backgroundColor: color
                }}
            >
                <Typography component="span">{header}</Typography>
            </AccordionSummary>
            <AccordionDetails>
                {details}
                {children}
                <br />
                <JSONPrinter dataObject={dataObject} />
            </AccordionDetails>
        </Accordion>
    )
}

export default Infrastructure;