import { User } from '../entities/User';
import { MyContext } from 'src/types';
import { Resolver, Arg, InputType, Field, Ctx, Mutation, ObjectType } from 'type-graphql';
import argon2 from 'argon2'


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
    @Mutation(() => UserResponse)
    async register(
        @Arg('options') options: UsernamePasswordInput,
        @Ctx() {em}: MyContext
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
        return {user,};
    }

    @Mutation(() => UserResponse)
    async login(
        @Arg('options') options: UsernamePasswordInput,
        @Ctx() {em}: MyContext
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
                        field: "crendentials",
                        message: "Invalid credentials",
                    },
                ],
            };
        }

        return {user,};
    }
}