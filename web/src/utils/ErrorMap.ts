import { FieldError } from "../generated/graphql";

export const toErrorMap = (error: FieldError) => {
    const errorMap: Record<string, string> = {};
    errorMap[error.field] = error.message
    return errorMap;
}