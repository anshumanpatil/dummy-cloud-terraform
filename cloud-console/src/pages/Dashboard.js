import { useEffect, useState } from 'react';
import Accordion from 'react-bootstrap/Accordion';
import Resource from "./resources/Resource";
import HTTPCalls from "../http/calls";


const ws = new WebSocket("ws://localhost:8888/ws");
// Triggered when the connection is opened
ws.onopen = function (evt) {
    console.log("Connection open...");
};
//Triggered when a message is received

//Triggered when the connection is closed
ws.onclose = function (evt) {
    console.log("Connection closed.", evt);
};

ws.onerror = function(err) {
    console.log("Connection error.", err);
}


export const Dashhboard = () => {
    const [instanceList, setInstanceList] = useState([]);
    const [bucketList, setBucketList] = useState([]);
    const [networkList, setNetworkList] = useState([]);

    useEffect(() => {
        HTTPCalls.requestAllInstances("").then(newinstanceList => setInstanceList(newinstanceList))
        HTTPCalls.requestAllBuckets("").then(newbucketList => setBucketList(newbucketList))
        HTTPCalls.requestAllNetworks("").then(newnetworkList => setNetworkList(newnetworkList))
    }, []);


    ws.onmessage = function (evt) {
        if (evt.data.includes("instance")) {
            HTTPCalls.requestAllInstances("").then(newinstanceList => setInstanceList(newinstanceList))
        }
        if (evt.data.includes("bucket")) {
            HTTPCalls.requestAllBuckets("").then(newbucketList => setBucketList(newbucketList))
        }
        if (evt.data.includes("network")) {
            HTTPCalls.requestAllNetworks("").then(newnetworkList => setNetworkList(newnetworkList))
        }

        console.log("Received Message: " + evt.data.includes("instance"));
    };

    return (
        <Accordion>
            <Accordion.Item eventKey="0">
                <Resource title="Instance" resourceList={instanceList} />
            </Accordion.Item>
            <Accordion.Item eventKey="1">
                <Resource title="Bucket" resourceList={bucketList}/>
            </Accordion.Item>
            <Accordion.Item eventKey="2">
                <Resource title="Network" resourceList={networkList}/>
            </Accordion.Item>
        </Accordion>
    );
}