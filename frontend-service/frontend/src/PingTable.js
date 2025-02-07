import React, { useEffect, useState } from 'react';
import { Table } from 'antd';

const PingTable = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);

    const columns = [
        {
            title: "Ip Address",
            dataIndex: "ipAddress",
            key: "ipAddress"
        },
        {
            title: "Ping Time",
            dataIndex: "pingTime",
            key: "pingTime"
        },
        {
            title: "Date last successful ping",
            dataIndex: "dateLastSuccessfulPing",
            key: "dateLastSuccessfulPing"
        },
    ];

    useEffect(() => {
        fetch('http://frontend:3001/data')
            .then(response => response.json())
            .then(data => {
                setData(data);
                setLoading(false);
            })
            .catch(error => {
                console.error("Error fetching ping results: ", error);
                setLoading(false);
            });

        }, []);

    return (
        <Table dataSource={data} columns={columns} loading={loading} rowKey="ipAddress" />
    );
};

export default PingTable;

