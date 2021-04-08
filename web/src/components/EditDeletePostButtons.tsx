import { EditIcon, DeleteIcon } from "@chakra-ui/icons";
import { Box, IconButton, Link } from "@chakra-ui/react";
import NextLink from "next/link";
import React from "react";
import { useDeletePostMutation, useMeQuery } from "../generated/graphql";

interface EditDeletePostButtonProps {
	id: number;
	creatorId: number;
}

export const EditDeletePostsButtons: React.FC<EditDeletePostButtonProps> = ({id, creatorId}) => {
	const [{data: meData}] = useMeQuery();
	const [,deletePost] = useDeletePostMutation();

	if (meData?.me?.id !== creatorId) {
		return null
	}

	return (
		<Box>
			<NextLink href="/post/edit/[id]" as={`/post/edit/${id}`}>
				<IconButton
					as={Link}
					mr={1}
					icon={<EditIcon />}
					aria-label="Update Post"
				/>
			</NextLink>
			<IconButton
				icon={<DeleteIcon />}
				aria-label="Delete Post"
				onClick={() => {
					deletePost({id: id});
				}}
			/>
		</Box> 
	);
}