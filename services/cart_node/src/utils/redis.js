import { createClient } from "redis";

export class Redis {
  static instance;
  client;

  constructor() {
    if (!Redis.instance) {
      this.client = createClient({
        url: `redis://${process.env.REDIS_HOST}:${process.env.REDIS_PORT}`,
      });
      this.client.connect().catch(console.error);
      Redis.instance = this;
    }
    return Redis.instance;
  }

  async set(key, value) {
    const _key = typeof key === "string" ? key : JSON.stringify(key);
    try {
      await this.client.set(_key, value, { EX: 60 * 60 * 24 });
    } catch (e) {
      console.error(e);
      return null;
    }
  }

  async get(key) {
    const _key = typeof key === "string" ? key : JSON.stringify(key);
    try {
      const reply = await this.client.get(_key);
      return reply;
    } catch (e) {
      console.error(e);
      return null;
    }
  }

  async del(key) {
    const _key = typeof key === "string" ? key : JSON.stringify(key);
    try {
      await this.client.del(_key);
    } catch (e) {
      console.error(e);
      return null;
    }
  }
}
