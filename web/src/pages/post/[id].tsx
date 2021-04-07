import { withUrqlClient } from "next-urql";
import React from "react";
import { createUrqlClient } from "../../utils/createUrqlClien";
import { useRouter } from "next/router";
import { usePostQuery } from "../../generated/graphql";
import { Layout } from "../../components/Layout";
import { Heading } from "@chakra-ui/layout";
import { Box } from "@chakra-ui/react";

const Post = ({}) => {
    const router = useRouter();
    const intId = typeof router.query.id === "string" ? parseInt(router.query.id) : -1
    const [{data, error, fetching}] = usePostQuery({
			pause: intId === -1,
			variables: {
					id: intId
			}
    });

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
				<Heading mb={4}>{data.post.title}</Heading>
				{data.post.text}
			</Layout>
    )
}

export default withUrqlClient(createUrqlClient, {ssr: true})(Post);