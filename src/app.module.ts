import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { RedisModule } from '@liaoliaots/nestjs-redis';

const options = {
  host: process.env.REDIS_HOST,
  port: parseInt(process.env.REDIS_PORT),
  db: parseInt(process.env.REDIS_DB),
};

@Module({
  imports: [RedisModule.forRoot({ config: options })],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
