# frozen_string_literal: true

RSpec.describe Websockets::UpdateBalanceService do
  describe '.call' do
    before { allow(CalculateBalanceService).to receive(:call).with('uah').and_return(241_516) }

    before { allow(ActionCable).to receive_message_chain(:server, :broadcast) }

    before { described_class.call('uah') }

    it 'broadcasts to NotificationChannel', :aggregate_failures do
      expect(ActionCable).to have_received(:server).at_least(1)

      expect(ActionCable.server)
        .to have_received(:broadcast)
        .with('NotificationsChannel', { type: :balance, value: '241 516.00' })
    end
  end
end
