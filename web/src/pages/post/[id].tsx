import { withUrqlClient } from "next-urql";
import React from "react";
import { createUrqlClient } from "../../utils/createUrqlClien";
import { Layout } from "../../components/Layout";
import { Flex, Heading } from "@chakra-ui/layout";
import { Box } from "@chakra-ui/react";
import { useGetPostFromUrl } from "../../utils/useGetPostFromUrl";
import { useRouter } from "next/router";
import { EditDeletePostsButtons } from "../../components/EditDeletePostButtons";

const Post = ({}) => {
	const router = useRouter();
	const [{data, error, fetching}] = useGetPostFromUrl(router);

		if (fetching) {
			return (
				<Layout>
					<div>loading...</div>
				</Layout>
			);
		}
		
		if(error) {
			return <div>{error.message}</div>
		}

		if (!data?.post) {
			return (
				<Layout>
					<Box>Could not find post</Box>
				</Layout>
			)
		}
		

    return (
			<Layout>
				<Flex justifyContent="space-between">
					<Heading mb={4}>{data.post.title}</Heading>
					<EditDeletePostsButtons 
						id={data.post.id} 
						creatorId={data.post.creator.id} 
					/>
				</Flex>
				{data.post.text}
			</Layout>
    )
}

export default withUrqlClient(createUrqlClient, {ssr: true})(Post);