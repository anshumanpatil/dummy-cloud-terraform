
import { Fragment } from 'react';
import Table from 'react-bootstrap/Table';
import Accordion from 'react-bootstrap/Accordion';


function Resource(props) {
    const { title, resourceList } = props
    const heads = resourceList && resourceList.length ? Object.keys(resourceList[0]) : []
    return (
        <>
            <Accordion.Header>{title}</Accordion.Header>
            <Accordion.Body>
                <Table striped bordered hover variant="dark" size="sm">
                
                    <thead>
                        <tr>
                            <th>#</th>
                            {heads.map( (e, i) => <th key={i}>{e}</th>)}
                        </tr>
                    </thead>
                    <tbody>
                        {resourceList && resourceList.length  && resourceList.map( (e, i) => 
                            <tr key={i}>
                                <td>{i}</td>
                                {heads.map( (et, it) => <td key={it}>{e[et]}</td>)}
                            </tr>)}
                        
                    </tbody>
                </Table>
            </Accordion.Body>
        </>
    );
}

export default Resource;