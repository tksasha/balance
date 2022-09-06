# frozen_string_literal: true

RSpec.describe Websockets::UpdateBalanceService do
  describe '.call' do
    before do
      allow(CalculateBalanceService).to receive(:call).with('uah').and_return(241_516)

      allow(ActionCable).to receive_message_chain(:server, :broadcast)

      described_class.call('uah')
    end

    it 'broadcasts to NotificationChannel', :aggregate_failures do
      expect(ActionCable).to have_received(:server).at_least(1)

      expect(ActionCable.server)
        .to have_received(:broadcast)
        .with('NotificationsChannel', { type: :balance, value: '241 516.00' })
    end
  end
end
