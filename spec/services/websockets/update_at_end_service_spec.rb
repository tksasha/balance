# frozen_string_literal: true

RSpec.describe Websockets::UpdateAtEndService do
  describe '.call' do
    before do
      allow(CalculateAtEndService).to receive(:call).with('uah').and_return(241_516)

      allow(ActionCable).to receive_message_chain(:server, :broadcast)

      described_class.call('uah')
    end

    it 'broadcasts to NotificationChannel', :aggregate_failures do
      expect(ActionCable).to have_received(:server)

      expect(ActionCable.server)
        .to have_received(:broadcast)
        .with('NotificationsChannel', { type: :at_end, value: '241 516.00' })
    end
  end
end
