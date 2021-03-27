import { User } from '../entities/User';
import { MyContext } from 'src/types';
import { Resolver, Arg, InputType, Field, Ctx, Mutation, ObjectType, Query } from 'type-graphql';
import argon2 from 'argon2';
//import {EntityManager} from '@mikro-orm/postgresql';


@InputType()
class UsernamePasswordInput {
    @Field()
    username: string

    @Field()
    password: string   
}

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
        if (options.username.length <= 2 ) {
            return {
                errors: [
                    {
                        field: "username",
                        message: "username must be at least 3 characters"
                    },
                ],
            };
        }

        if (options.password.length <= 4) {
            return {
                errors: [
                    {
                        field: "password",
                        message: "password mut be at least 4 characters"
                    },
                ],
            };
        }
        const hashedPassword = await argon2.hash(options.password);
        const user = em.create(User, {
            username: options.username, 
            password: hashedPassword,
        });

        try {
            // (em as EntityManager).createQueryBuilder(User).getKnexQuery().insert({
            //         username: options.username, 
            //         password: hashedPassword,
            //         createdAt: new Date(),
            //         updatedAt: new Date()
            // });
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
        @Arg('options') options: UsernamePasswordInput,
        @Ctx() { em, req }: MyContext
    ): Promise<UserResponse> {
        const user = await em.findOne(User, {username: options.username});
        if (!user) {
            return {
                errors: [
                    {
                        field: 'username',
                        message: "that username doesn't exist"
                    }, 
                ],
            };
        }
        const valid = await argon2.verify(user.password, options.password);
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
}