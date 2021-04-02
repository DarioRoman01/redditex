import React from 'react';
import {Form, Formik} from 'formik';
import { Box, Button, Flex, Link } from '@chakra-ui/react';
import { Wrapper } from '../components/Wrapper';
import { InputField } from '../components/InputField';
import { useLoginMutation } from '../generated/graphql';
import { toErrorMap } from '../utils/ErrorMap';
import { useRouter } from 'next/router';
import { withUrqlClient } from 'next-urql';
import { createUrqlClient } from '../utils/createUrqlClien';
import NextLink from 'next/link';

const Login: React.FC<{}> = ({}) => {
  const router = useRouter();
  const [,login] = useLoginMutation();
    return (
        <Wrapper variant="small">
          <Formik 
            initialValues={{usernameOrEmail: "", password: ""}} 
            onSubmit={async (values, {setErrors}) => {
                const response =  await login(values);
                if (response.data?.login.error) {
                  setErrors(toErrorMap(response.data.login.error));
                } else if (response.data?.login.user){
                  router.push("/");
                }
            }}
            >
            {({isSubmitting}) => (
              <Form>
                <InputField 
                  name="usernameOrEmail" 
                  placeholder="username or email" 
                  label="Username or Email"
                />
                <Box mt={4}>
                  <InputField 
                    name="password" 
                    placeholder="password" 
                    label="Password" 
                    type="password"
                  />
                </Box>
                <Flex mt={2}>
                  <NextLink href="/forgot-password">
                    <Link ml="auto">forgot password?</Link>
                  </NextLink>
                </Flex>
                <Button
                  mt={4} 
                  type="submit"
                  isLoading={isSubmitting} 
                  colorScheme="teal"
                >
                  Login
                </Button>
              </Form>
            )}
          </Formik>
        </Wrapper>
    );
};

export default withUrqlClient(createUrqlClient)(Login);