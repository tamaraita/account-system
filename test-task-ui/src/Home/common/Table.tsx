import React from 'react';
import { Table as AnTable,
    Typography } from 'antd';

import { Transaction } from '../../lib/ducks/transactions';


const { Text } = Typography;

interface TableProps {
    data: Transaction[]
}

const Table = ({ data }: TableProps): React.ReactElement => {

    const renderDescription = (data: Transaction): React.ReactElement => {
        return (
            <div style={{display:"flex", flexDirection:"column"}}>
                <Text strong>ID:</Text><Text>{data.id}</Text>
                <Text strong>Type:</Text><Text>{data.type}</Text>
                <Text strong>Amount:</Text><Text>${data.amount}</Text>
                <Text strong>Effective date:</Text><Text>{data.effectiveDate}</Text>
            </div>
        )
    }

    const createKeys = (record: any): string => {
        return `row-${record.id}`;
    };

    const columns = [
        {
            title: 'Type',
            dataIndex: 'type',
            render: (ttype: string) => <Text style={{color:ttype.includes("debit")? "green": "red"}}>{ttype}</Text>
        },
        {
            title: 'Amount',
            dataIndex: 'amount',
            render: (amount: number) => <Text>${amount}</Text>,
        }
    ]

    return (
        <AnTable
            columns={columns}
            dataSource={data}
            expandable={{
                expandedRowRender: (record: Transaction) => renderDescription(record),
                expandRowByClick: true,
            }}
            rowKey={createKeys}
        />
    );
};

export default Table;