import { useEffect, useState } from 'react';
import Accordion from 'react-bootstrap/Accordion';
import Resource from "./resources/Resource";
import HTTPCalls from "../http/calls";

export const Dashhboard = () => {
    const [instanceList, setInstanceList] = useState([]);
    const [bucketList, setBucketList] = useState([]);

    useEffect(() => {
        HTTPCalls.requestAllInstances("").then(newinstanceList => setInstanceList(newinstanceList))
        HTTPCalls.requestAllBuckets("").then(newbucketList => setBucketList(newbucketList))
    }, []);

    const accChanged = async (e) => {
        if(!e) return;
        switch (parseInt(e)) {
            case 0:
                HTTPCalls.requestAllInstances("").then(newinstanceList => setInstanceList(newinstanceList))
                break;
            case 1:
                HTTPCalls.requestAllBuckets("").then(newbucketList => setBucketList(newbucketList))
                break;
        
            default:
                break;
        }
        
        
    }
    return (
        <Accordion onSelect={accChanged}>
            <Accordion.Item eventKey="0">
                <Resource title="Instance" resourceList={instanceList} />
            </Accordion.Item>
            <Accordion.Item eventKey="1">
                <Resource title="Bucket" resourceList={bucketList}/>
            </Accordion.Item>
        </Accordion>
    );
}