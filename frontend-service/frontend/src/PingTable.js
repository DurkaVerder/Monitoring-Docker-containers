import React, { useEffect, useState } from 'react';
import { Table } from 'antd';
import moment from 'moment';

const PingTable = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);

    const columns = [
        {
            title: "Ip Address",
            dataIndex: "ip_address",
            key: "ip_address"
        },
        {
            title: "Ping Time",
            dataIndex: "ping_time",
            key: "ping_time"
        },
        {
            title: "Date last successful ping",
            key: "date_successful_ping",
            render: (_, record) => (
              record.date_successful_ping.Valid 
                ? moment(record.date_successful_ping.Time).format("YYYY-MM-DD HH:mm:ss") 
                : "N/A"
            )
        },
    ];

    useEffect(() => {
        const fetchData = () => {
            fetch('http://localhost:3001/data')
                .then(response => response.json())
                .then(data => {
                    setData(data);
                    setLoading(false);
                })
                .catch(error => {
                    console.error("Error fetching ping results: ", error);
                    setLoading(false);
                });
        };

        fetchData();

        const intervalId = setInterval(fetchData, 5000);

    
        return () => clearInterval(intervalId);
    }, []);

    return (
        <Table 
            dataSource={data} 
            columns={columns} 
            loading={loading} 
            rowKey="id" 
        />
    );
};

export default PingTable;