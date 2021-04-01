import { Box, Button } from '@chakra-ui/react';
import { Formik, Form } from 'formik';
import { NextPage } from 'next';
import { withUrqlClient } from 'next-urql';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import { InputField } from '../../components/InputField';
import { Wrapper } from '../../components/Wrapper';
import { useChangePasswordMutation } from '../../generated/graphql';
import { createUrqlClient } from '../../utils/createUrqlClien';
import { toErrorMap } from '../../utils/ErrorMap';

const ChangePassword: NextPage<{token: string}> = ({ token }) => {
  const router = useRouter();
  const [,changePassword] = useChangePasswordMutation();
  const [tokenError, setTokenError] = useState("");
  return (
    <Wrapper variant="small">
    <Formik 
      initialValues={{newPassword: ""}} 
      onSubmit={async (values, {setErrors}) => {
        const response =  await changePassword({
          newPassword: values.newPassword, 
          token: token
        });

        if (response.data?.changePassword.error) {
          const errorMap = toErrorMap(response.data.changePassword.error)
          if ('token' in errorMap) {
            setTokenError(errorMap.token);
          }
          setErrors(errorMap);
        } else if (response.data?.changePassword.user){
            router.push("/");
          }
      }}
      >
      {({isSubmitting}) => (
        <Form>
          <InputField 
            name="newPassword" 
            placeholder="new password" 
            label="New Password"
            type="password"
          />
          {tokenError ? <Box color="red">{tokenError}</Box> : null}
          <Button
            mt={4}
            type="submit"
            isLoading={isSubmitting}
            colorScheme="teal"
          >
            change password
          </Button>
        </Form>
        
      )}
  </Formik>
  </Wrapper>
  );
}

ChangePassword.getInitialProps = ({query}) => {
  return {
    token: query.token as string,
  };
};

export default withUrqlClient(createUrqlClient)(ChangePassword as any);