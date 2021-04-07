import { Box, Button, Flex, Heading, Link } from '@chakra-ui/react';
import React from 'react';
import NextLink from 'next/link';
import { useLogoutMutation, useMeQuery } from '../generated/graphql';
import { isServer } from '../utils/isServer';

interface NavBarProps {}

export const NavBar: React.FC<NavBarProps> = ({}) => {
  const [{fetching: LogoutFetching},logout] = useLogoutMutation()
    const [{data, fetching}] = useMeQuery({
      pause: isServer(),
    });
    let body = null 
      
    // data is loading
    if (fetching) {
      // user not logged in
    } else if (!data?.me) {
      body = (
        <>
        <NextLink href="/login">
          <Link 
            color="blackAlpha.900"
            fontSize="larger"
            mr={2} 
          >
            Login
          </Link>
        </NextLink>
        <NextLink href="/register">
          <Link 
            color="blackAlpha.900" 
            fontSize="larger"
          >
            Register
          </Link>
        </NextLink>
        </>
      );
      // user is logged in
    } else {
        body = (
          <Flex alignItems="center">
            <NextLink href="/create-post">
              <Button
                as={Link}
                ml="auto" 
                mr={3}
                fontSize="large"
                color="blackAlpha.900"
                colorScheme="blackAlpha" 
                variant="outline"
              >
                  Create Post
              </Button>
            </NextLink>
            <Box mr={3} fontSize="larger">{data.me.username}</Box>
            <Link
              color="blackAlpha.900"
              fontSize="larger"
              onClick={() => {
                logout();
              }}
              isLoading={LogoutFetching}
            >
              Logout
            </Link>
          </Flex>
        );
    }

    return (
      <Flex zIndex={1} position="sticky" top={0} bg="twitter.400" p={4} alignItems="center">
        <NextLink href="/">
          <Link>
            <Heading>Redditex</Heading>
          </Link>
        </NextLink>
          <Box ml={"auto"}>{body}</Box>
      </Flex>
    );
}