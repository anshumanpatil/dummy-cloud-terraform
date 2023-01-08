
import { Fragment } from 'react';
import Table from 'react-bootstrap/Table';
import Accordion from 'react-bootstrap/Accordion';


function Resource(props) {
    let { title, resourceList } = props
    const heads = resourceList && resourceList.length ? Object.keys(resourceList[0]) : []
    if(heads.length && heads.indexOf("iplist") >= 0) {
        // heads.map(e => console.log(e))
        resourceList = resourceList.map(e => {
            return {
                ...e,
                instancelist: JSON.stringify(e.instancelist),
                iplist: JSON.stringify(e.iplist),
                isactive: (e.isactive ? "TRUE" : "FALSE")
            }
        })
    }

    return (
        <>
            <Accordion.Header>{title}</Accordion.Header>
            <Accordion.Body>
                <Table striped bordered hover size="sm">
                
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