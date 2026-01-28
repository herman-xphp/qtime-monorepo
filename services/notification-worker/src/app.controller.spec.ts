import { Test, TestingModule } from '@nestjs/testing';
import { AppController } from './app.controller';

describe('AppController', () => {
  let appController: AppController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [AppController],
    }).compile();

    appController = app.get<AppController>(AppController);
  });

  describe('handleTicketCreated', () => {
    it('should log the event', () => {
      // Mock Console
      const logSpy = jest.spyOn(console, 'log').mockImplementation();

      const payload = { merchant_id: 'UnitTester', ticket_number: 99 };
      appController.handleTicketCreated(payload);

      expect(logSpy).toHaveBeenCalled();

      // Cleanup
      logSpy.mockRestore();
    });
  });
});
