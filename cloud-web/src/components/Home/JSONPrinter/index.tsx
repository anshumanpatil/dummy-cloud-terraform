
import * as React from 'react';
import { PropsWithChildren } from 'react';

interface JSONPrinterBaseProps {
    dataObject: any;
}
type JSONPrinterProps = PropsWithChildren<JSONPrinterBaseProps>;

const JSONPrinter: React.FC<JSONPrinterProps> = ({ dataObject }) => {
    return (
        <div style={{ backgroundColor: 'black', padding: '16px', borderRadius: '2px', color: 'white', width: '50%', overflow: 'auto' }}>
            <pre>
                {JSON.stringify(dataObject, null, 4)}
            </pre>
        </div>
    )
}
export default JSONPrinter;