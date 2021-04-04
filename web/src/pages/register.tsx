import React from 'react';
import {Form, Formik} from 'formik';
import { Box, Button } from '@chakra-ui/react';
import { Wrapper } from '../components/Wrapper';
import { InputField } from '../components/InputField';
import { useRegisterMutation } from '../generated/graphql';
import { toErrorMap } from '../utils/ErrorMap';
import { useRouter } from 'next/router';
import { withUrqlClient } from 'next-urql';
import { createUrqlClient } from '../utils/createUrqlClien';

interface registerProps {}


const Register: React.FC<registerProps> = ({}) => {
  const router = useRouter();
  const [,register] = useRegisterMutation();
    return (
        <Wrapper variant="small">
          <Formik 
            initialValues={{username: "", email: "",password: ""}} 
            onSubmit={async (values, {setErrors}) => {
                const response =  await register({options: values});
                if (response.data?.register.error) {
                  setErrors(toErrorMap(response.data.register.error));
                } else if (response.data?.register.user){
                  router.push("/");
                }
            }}
            >
            {({isSubmitting}) => (
              <Form>
                <InputField 
                  name="username" 
                  placeholder="username" 
                  label="Username"
                />
                <Box mt={4}>
                  <InputField 
                    name="email" 
                    placeholder="email" 
                    label="Email" 
                  />
                </Box>
                <Box mt={4}>
                  <InputField 
                    name="password" 
                    placeholder="password" 
                    label="Password" 
                    type="password"
                  />
                </Box>
                <Button 
                  mt={4} 
                  type="submit"
                  isLoading={isSubmitting} 
                  colorScheme="twitter"
                >
                  Register
                </Button>
              </Form>
            )}
        </Formik>
        </Wrapper>
       
    );
};

export default withUrqlClient(createUrqlClient)(Register);