import pkg from "pg";
const { Pool } = pkg;

import dotenv from "dotenv";
dotenv.config();

export class Database {
  constructor() {
    if (Database.instance) {
      return Database.instance;
    }
    Database.instance = this;
    this.pool = new Pool({
      host: process.env.DB_HOST,
      port: process.env.DB_PORT,
      user: process.env.DB_USER,
      database: process.env.DB_NAME,
      password: process.env.DB_PASS,
      max: 20,
      idleTimeoutMillis: 30000,
      connectionTimeoutMillis: 2000,
    });
  }

  async query(query, values) {
    console.log(`Executing query: ${query}`);
    console.log(`Query values: ${JSON.stringify(values)}`);
    const client = await this.pool.connect();
    try {
      return await client.query(query, values);
    } finally {
      client.release();
    }
  }
}
