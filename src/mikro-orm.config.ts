import { MikroORM } from "@mikro-orm/core";
import { __prod__ } from "./constants";
import { Post } from "./entities/Posts";
import path from 'path';
import { User } from "./entities/User";

// Mikro-orm dev config
export default {
    migrations:{
        path: path.join(__dirname, "./migrations"),
        pattern: /^[\w-]+\d+\.[tj]s$/,
    },
    entities: [Post, User],
    clientUrl: "http://localhost:5432",
    dbName: "lireddit",
    user: "postgres",
    password: "admin123",
    type: "postgresql",
    debug: !__prod__,
} as Parameters<typeof MikroORM.init>[0];