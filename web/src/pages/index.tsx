import { Box, Button, Flex, Heading, Link, Stack, Text } from "@chakra-ui/react";
import { withUrqlClient } from 'next-urql';
import NextLink from "next/link";
import React, { useState } from 'react';
import { EditDeletePostsButtons } from "../components/EditDeletePostButtons";
import { Layout } from "../components/Layout";
import { UpdootSection } from "../components/UpdootSection";
import { usePostsQuery } from "../generated/graphql";
import { createUrqlClient } from "../utils/createUrqlClien";

const Index = () => {
  const [variables, setVariables] = useState({
    limit: 15, 
    cursor: null as null | string,
  });

  const [{ data, fetching }] = usePostsQuery({
    variables,
  });

  if (!fetching && !data) {
    return (
      <Layout>
        <div>you got query error for some reason</div>
      </Layout>
    );
  }

  return(
    <Layout>
      {!data && fetching ? (
        <div>loading....</div>
      ) : (
        <Stack spacing={8}>
          {data!.posts.posts.map((p) => !p ? null : (
            <Flex key={p.id} p={5} shadow="md" borderWidth="1px">
              <UpdootSection post={p} />
              <Box flex={1}>
                <NextLink href="/post/[id]" as={`/post/${p.id}`}>
                  <Link>
                    <Heading fontSize="xl">{p.title}</Heading>
                  </Link>
                </NextLink>
                <Text>Posted By {p.creator.username}</Text> 
                <Flex aling="center">
                  <Text flex={1} mt={4}>{p.textSnippet}</Text>
                    <Box ml="auto">
                      <EditDeletePostsButtons id={p.id} creatorId={p.creator.id} />
                    </Box>
                </Flex>
              </Box>
          </Flex>
          ))}
        </Stack>
        )}
        {data && data.posts.hasMore ? (
          <Flex >
            <Button
              onClick={() => {
                setVariables({
                  limit: variables.limit,
                  cursor: data.posts.posts[data.posts.posts.length - 1].createdAt,
                })
              }}
              m="auto"
              my={8}
              isLoading={fetching}
              colorScheme="twitter"
            >
              load more
            </Button>
          </Flex>
        ) : null}
    </Layout>
  );
}

export default withUrqlClient(createUrqlClient, {ssr: true})(Index);
