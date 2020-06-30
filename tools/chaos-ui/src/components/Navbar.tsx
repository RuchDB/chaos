import { Box, Flex, Image, Link, PseudoBox, Stack, useColorMode, IconButton, Heading } from "@chakra-ui/core";
import React from "react";


const Navbar: React.SFC = () => {
    const { colorMode, toggleColorMode } = useColorMode();
    const bgColor = { light: 'gray.300', dark: 'gray.600' };
    const textColor = { light: 'black', dark: 'gray.100' };
    // const router = useRouter();

    return (
        <Flex
            w='100vw'
            bg={bgColor[colorMode]}
            align='center'
            color={textColor[colorMode]}
            justify='center'
            fontSize={['md', 'lg', 'xl', 'xl']}
            h='7vh'
            boxShadow='md'
            p={2}>
            <Flex w={['100vw', '100vw', '80vw', '80vw']} justify='space-around'>
                <Box>
                    <Image h='4vh' src='./logo.svg' alt='Logo of Chakra-ui' />
                </Box>
                <Stack
                    spacing={8}
                    color={textColor[colorMode]}
                    justify='center'
                    align='center'
                    isInline>
                    <PseudoBox
                        position='relative'
                        opacity={0.4}>
                        <Heading as="h3" size="lg">
                            Chaos Project
                            </Heading>
                    </PseudoBox>
                    {/* <PseudoBox
                        position='relative'
                        opacity={0.4}>
                        <Link href='/form'>
                            <a>Form</a>
                        </Link>
                    </PseudoBox>
                    <PseudoBox
                        position='relative'
                        opacity={0.4}>
                        <Link href='/card'>
                            <a>Card</a>
                        </Link>
                    </PseudoBox>
                    <PseudoBox
                        position='relative'
                        opacity={0.4}>
                        <Link href='/list'>
                            <a>List</a>
                        </Link>
                    </PseudoBox> */}
                </Stack>
                <Box>
                    <IconButton
                        aria-label=''
                        rounded='full'
                        onClick={toggleColorMode}
                        icon={colorMode === 'light' ? 'moon' : 'sun'}>
                        Change Color Mode
					</IconButton>
                </Box>
            </Flex>
        </Flex>
    )
}


export default Navbar;