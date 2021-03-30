import { User } from '../entities/User';
import { MyContext } from 'src/types';
import { Resolver, Arg, Field, Ctx, Mutation, ObjectType, Query } from 'type-graphql';
import argon2 from 'argon2';
import { COOKIE_NAME } from '../constants';
import { UsernamePasswordInput } from './UsernamePasswordInput';
import { validateRegister } from '../utils/validateRegister';

@ObjectType()
class FieldError {
    @Field()
    field: string;

    @Field()
    message: string;
}

@ObjectType()
class UserResponse {
    @Field(() => [FieldError], {nullable: true})
    errors?: FieldError[];

    @Field(() => User, {nullable: true})
    user?: User;
}

@Resolver()
export class UserResolver {
    @Mutation(() => Boolean)
    async forgotPassword(
        @Arg("email") email: string,
        @Ctx() {em}: MyContext
    ) {
        const person = await em.findOne(User, {email: email})
        if (!person) {
          return {
            errors: [
              {
                field: "email",
                message: "Email not registered"
              }, 
            ],
          };
        }
        return true;
    }   

    // return user info if its logged in
    @Query(() => User, {nullable: true})
    async me(
        @Ctx() { req, em }: MyContext
    ) {
        // you are not logged in
        if (!req.session.userId) {
            return null;
        }

        const user = await em.findOne(User, {id: req.session.userId});
        return user;
    }


    // register mutation handle validation and store user data in the db
    @Mutation(() => UserResponse)
    async register(
        @Arg('options') options: UsernamePasswordInput,
        @Ctx() { em, req }: MyContext
    ): Promise<UserResponse> {
        const errors = validateRegister(options);
        if (errors) { 
            return {errors}; 
        }

        const hashedPassword = await argon2.hash(options.password);
        const user = em.create(User, {
            username: options.username, 
            email: options.email,
            password: hashedPassword,
        });

        try {
            await em.persistAndFlush(user); 
        } catch(err) {
            if (err.code === '23505') { // || err.detail.includes("lready exists")
              return {
                errors: [
                  {
                    field: "username",
                    message: "That username is already in use"
                  },
                ],
              };
            }
        }

        // store user id session
        // this will set a cookie on ther user
        // keep them logged in
        req.session.userId = user.id
        return {user,};
    }


    // Login mutation verify credentials and send a cokie to the client
    @Mutation(() => UserResponse)
    async login(
        @Arg('usernameOrEmail') usernameOrEmail: string,
        @Arg("password") password: string,
        @Ctx() { em, req }: MyContext
    ): Promise<UserResponse> {
        const user = await em.findOne(User, 
            usernameOrEmail.includes("@")
            ? {email: usernameOrEmail}  
            : { username: usernameOrEmail}
        );

        if (!user) {
          return {
            errors: [
              {
                field: 'usernameOrEmail',
                message: "that username doesn't exist"
              }, 
            ],
          };
        }

        const valid = await argon2.verify(user.password, password);
        if (!valid) {
          return {
            errors: [
              {
                field: "password",
                message: "Invalid credentials",
              },
            ],
          };
        }

        req.session.userId = user.id;

        return {user,};
    }

    @Mutation(() => Boolean)
    logout(@Ctx() { req, res }: MyContext) {
        return new Promise((resolve) =>
            req.session.destroy(err => {
                res.clearCookie(COOKIE_NAME);
                if (err) {
                    console.log(err);
                    resolve(false);
                    return;
                }

                resolve(true);
            })
        );
    }
}