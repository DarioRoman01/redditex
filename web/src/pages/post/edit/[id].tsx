import { Box, Button } from "@chakra-ui/react";
import { Form, Formik } from "formik";
import { withUrqlClient } from "next-urql";
import { useRouter } from "next/router";
import React from "react";
import { InputField } from "../../../components/InputField";
import { Layout } from "../../../components/Layout";
import { useUpdatePostMutation } from "../../../generated/graphql";
import { createUrqlClient } from "../../../utils/createUrqlClien";
import { useGetPostFromUrl } from "../../../utils/useGetPostFromUrl";

const EditPost  = ({}) => {
	const router = useRouter();
	const [{data, fetching}] = useGetPostFromUrl(router);
	const [,updatePost] = useUpdatePostMutation();
	if (fetching) {
		return (
			<Layout>
				<div>loading...</div>
			</Layout>
		);
	}

	if (!data?.post) {
		return (
			<Layout>
				<div>could not find post</div>
			</Layout>
		);
	}

	return (
		<Layout variant="small">
		<Formik 
			initialValues={{title: data.post.title, text: data.post.text}} 
			onSubmit={async (values) => {
				await updatePost({id: data.post.id, options: values});
				router.back();
			}}
		>
			{({isSubmitting}) => (
				<Form>
					<InputField 
						name="title" 
						placeholder="title" 
						label="Title"
					/>
					<Box mt={4}>
						<InputField 
							textarea
							name="text" 
							placeholder="text..." 
							label="Body"
						/>
					</Box>
					<Button
						mt={4} 
						type="submit"
						isLoading={isSubmitting} 
						colorScheme="twitter"
					>
						Update Post
					</Button>
				</Form>
			)}
		</Formik>
	</Layout>
	);
}

export default withUrqlClient(createUrqlClient)(EditPost);