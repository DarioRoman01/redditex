import { UsernamePasswordInput } from "src/resolvers/UsernamePasswordInput";

export const validateRegister = (options: UsernamePasswordInput) => {
    if (!options.email.includes("@")) {
        return [
          {
            field: "username",
            message: "Invalid email"
          },
        ]
      }

      if (options.username.length <= 2 ) {
        return [
          {
            field: "username",
            message: "username must be at least 3 characters"
          },
        ]
      }

      if (options.username.includes("@")) {
        return [
          {
            field: "username",
            message: "cannot include an @"
          },
        ]
      }


      if (options.password.length <= 4) {
        return [
          {
            field: "password",
            message: "password mut be at least 4 characters"
          },
        ]
      }

    return null;
}