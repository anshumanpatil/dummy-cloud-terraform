import { Navigate } from "react-router-dom";
import { Fragment } from "react";
import { Dashhboard } from "./Dashboard";
import { LocalStorage } from "../Auth/LocalStorage";

import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

export const HomePage = () => {
    const { user } = LocalStorage.get("user");
        if (!user) {
            window.location.href = "/"
        }
    
    return (
        <Container>
                <div style={{height:'1rem'}}></div>

            <Row>
                <Dashhboard></Dashhboard>
            </Row>
        </Container>
    );
}
