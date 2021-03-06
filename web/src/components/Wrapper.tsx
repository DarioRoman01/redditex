import { Box } from '@chakra-ui/react';
import React from 'react'

export type WrapperVariant = "small" | "regular"

interface WraperProps {
    variant?: WrapperVariant;
}

export const Wrapper: React.FC<WraperProps> = ({children, variant='regular'}) => {
    return (
      <Box 
        mt={8} 
        mx="auto" 
        maxW={variant === "regular" ? "800px" : "400px"} 
        w="100%"
      >
        {children}
      </Box>
    );   
}