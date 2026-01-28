import { Controller } from '@nestjs/common';
import { EventPattern, Payload } from '@nestjs/microservices';

interface TicketEvent {
  merchant_id: string;
  ticket_number: number;
}

@Controller()
export class AppController {
  @EventPattern('queue-events')
  handleTicketCreated(@Payload() data: TicketEvent) {
    // In NestJS Kafka, the payload behaves differently depending on serializer.
    // Usually 'data' is the 'value' of the message.
    console.log('🔔 [Notification Worker] Received Event:', data);

    // Simulate Sending WhatsApp
    const event = data;
    if (event.merchant_id && event.ticket_number) {
      console.log(
        `📲 Sending WA to Merchant ${event.merchant_id}: "New Ticket #${event.ticket_number} created!"`,
      );
    }
  }
}
