import { NestFactory } from '@nestjs/core';
import { Transport, MicroserviceOptions } from '@nestjs/microservices';
import { AppModule } from './app.module';
import { CompressionTypes, CompressionCodecs } from 'kafkajs';
// eslint-disable-next-line @typescript-eslint/no-require-imports, @typescript-eslint/no-unsafe-assignment
const SnappyCodec = require('kafkajs-snappy');

// eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
CompressionCodecs[CompressionTypes.Snappy] = SnappyCodec;

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    AppModule,
    {
      transport: Transport.KAFKA,
      options: {
        client: {
          brokers: (process.env.KAFKA_BROKERS || 'localhost:9092').split(','),
        },
        consumer: {
          groupId: 'notification-worker',
        },
      },
    },
  );
  await app.listen();
}
bootstrap().catch((err) => {
  console.error('Bootstrap error:', err);
});
