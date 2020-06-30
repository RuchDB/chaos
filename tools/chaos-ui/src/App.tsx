import { Avatar, Box, CSSReset, Flex, Grid, Stack } from '@chakra-ui/core';
import React from 'react';
import './App.css';
import CodeBar from './components/CodeBar';
import Navbar from './components/Navbar';


function App() {
  return (
    <Flex direction='column' align='center' justify='center'>
      <CSSReset />
      <Navbar />
      <Grid templateColumns="repeat(5, 1fr)" gap={6}>
        <Box w="100%" h="10" bg="blue.500" />
        <Box w="100%" h="10" bg="blue.500" />
        <Box w="100%" h="10" bg="blue.500" />
        <Box w="100%" h="10" bg="blue.500" />
        <Box w="100%" h="10" bg="blue.500" />
      </Grid>
      <Stack isInline>
        <Avatar size="xl" name="Christian Nwamba" src="https://avatars3.githubusercontent.com/u/15232997?s=460&u=11d9a6a8375c45936f1e3eb9f101fc24e3ee60ce&v=4" />
        <Avatar size="xl" name="Christian Nwamba" src="https://avatars0.githubusercontent.com/u/60535886?s=460&u=a7882f72d00b6239887a18a411cf9355d291c085&v=4" />
        <Avatar size="xl" name="Christian Nwamba" src="https://avatars1.githubusercontent.com/u/6275168?s=460&u=437e902ac95d088d783ede0a40e5e6c70be02665&v=4" />
      </Stack>
      <CodeBar />
    </Flex>
  )
}

export default App;
