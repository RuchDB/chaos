import { Button, Flex, CSSReset } from '@chakra-ui/core';
import React from 'react';
import './App.css';
import Navbar from './components/Navbar';
import CodeBar from './components/CodeBar';


function App() {
  return (
    <Flex direction='column' align='center' justify='center'>
      <CSSReset />
      <Navbar />
      <Flex justify='center' align='center' w='100%' h='93vh'>
        {/* <Component {...pageProps} /> */}
      </Flex>
    </Flex>
  )
}

export default App;
