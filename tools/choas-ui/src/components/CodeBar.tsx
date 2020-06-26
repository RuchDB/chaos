import { Box, Stat, StatHelpText, StatLabel, StatNumber } from '@chakra-ui/core';
import React from 'react';
import { Bar, BarChart, CartesianGrid, Legend, Tooltip, XAxis } from 'recharts';
import axios from 'axios';


class CodeBar extends React.Component {
    state: any = {
        data: [{ name: 'Code Stats', peterhp: 0, littleroys: 0, Jonir: 0 }],
        total: 0
    }

    componentDidMount() {
        axios.get('http://localhost:5000/stats').then(
            rsp => {
                this.setState({ total: rsp.data.total })
                this.setState({
                    data: [
                        {
                            name: 'Code Stats',
                            peterhp: rsp.data.stats.peterhp,
                            littleroys: rsp.data.stats.littleroys,
                            Jonir: rsp.data.stats.Jonir
                        }
                    ]
                })
            }, err => {
                console.log(err)
            }
        )
    }

    render() {
        return (
            <>
                <BarChart
                    width={600}
                    height={400}
                    data={this.state.data}
                    margin={{
                        top: 5, right: 30, left: 20, bottom: 5,
                    }}
                >
                    <CartesianGrid strokeDasharray="1 1" />
                    <XAxis dataKey="name" />
                    <XAxis />
                    <Tooltip />
                    <Legend />
                    <Bar dataKey="peterhp" fill="#8884d8" />
                    <Bar dataKey="littleroys" fill="#82ca9d" />
                    <Bar dataKey="Jonir" fill="#d09e2f" />
                </BarChart>
                <Box p={5} shadow="md" borderWidth="1px" >
                    <Stat>
                        <StatLabel>Total Line</StatLabel>
                        <StatNumber>{this.state.total}</StatNumber>
                        <StatHelpText>{Date()}</StatHelpText>
                    </Stat>
                </Box>
            </>
        )
    }
}

export default CodeBar;