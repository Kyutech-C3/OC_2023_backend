import { RedisService } from '@liaoliaots/nestjs-redis';
import { Injectable } from '@nestjs/common';
import { Redis } from 'ioredis';

@Injectable()
export class AppService {
  private readonly redis: Redis;
  constructor(private readonly redisService: RedisService) {
    this.redis = this.redisService.getClient();
  }

  async set(key: string, value: string) {
    return await this.redis.set(key, value);
  }

  getHello(): string {
    return 'Hello World!';
  }
}
